# sync.Pool

sync.Pool 提供了对象池的功能。它对外提供了三个方法：New、Get 和 Put。

## 源码分析

```go
type Pool struct {
    noCopy noCopy // 代表结构体禁止拷贝，使用`go vet`时生效
    local     unsafe.Pointer // poolLocal数组的指针
    localSize uintptr // 数组大小
    victim     unsafe.Pointer // victim也是一个poolLocal数组的指针，每次垃圾回收的时候，Pool 会把 victim 中的对象移除，然后把 local 的数据给 victim
    victimSize uintptr
    New func() interface{} // New函数是在创建pool的时候设置的，当pool没有缓存对象的时候，会调用New方法生成一个新的对象。
}
```

![pic](https://img.luozhiyun.com/20201226184341.png)

```go
type poolLocal struct {
    poolLocalInternal
    pad [128 - unsafe.Sizeof(poolLocalInternal{}) % 128]byte
}
```

poolLocal 数组的大小时 goroutine 中 P 的数量，访问时，P 的 id 对应 poolLocal 数组下标索引，所以 Pool 的最大个数是`runtime.GOMAXPROCS(0)`。通过这样的设计，每个 P
都有了自己的本地空间，多个 goroutine 使用同一个 Pool 时，减少了竞争，提升了性能。

poolLocal 里面有一个 pad 数组用来占位用，防止在 cache line 上分配多个 poolLocalInternal 从而造成 false
sharing，参考[What’s false sharing and how to solve it](https://medium.com/@genchilu/whats-false-sharing-and-how-to-solve-it-using-golang-as-example-ef978a305e10)

```go
type poolLocalInternal struct {
    private interface{} // private代表缓存的一个元素，只能由相应的一个 P 存取。(因为一个 P 同时只能执行一个 goroutine，所以不会有并发的问题)
    shared poolChain // Local P can pushHead/popHead; any P can popTail.
}
```

```go
type poolChain struct {
    head *poolChainElt
    tail *poolChainElt
}

type poolChainElt struct {
    poolDequeue
    next, prev *poolChainElt
}

type poolDequeue struct {
    headTail uint64
    vals []eface
}
```

poolChain 是一个双端队列，里面的 head 和 tail 分别指向队列头尾；poolDequeue 里面存放真正的数据，是一个单生产者、多消费者的固定大小的无锁的环状队列，headTail
是环状队列的首位位置的指针，可以通过位运算解析出首尾的位置，生产者可以从 head 插入、head 删除，而消费者仅可从 tail 删除。

![pic](https://img.luozhiyun.com/20201226184348.png)

## Get 方法

```go
func (p *Pool) Get() interface{} {
 if race.Enabled {
  race.Disable()
 }
 // 把当前goroutine绑定在当前的P上
 l, pid := p.pin()
 // 优先从local的private中获取数据
 x := l.private
 l.private = nil
 if x == nil {
  // Try to pop the head of the local shard. We prefer
  // the head over the tail for temporal locality of
  // reuse.
  x, _ = l.shared.popHead()
  // 如果shared的head也不存在，去其他P的local上偷一个
  if x == nil {
   x = p.getSlow(pid)
  }
 }
 // 解除抢占
 runtime_procUnpin()
 if race.Enabled {
  race.Enable()
  if x != nil {
   race.Acquire(poolRaceAddr(x))
  }
 }
 // 如果还是nil，调用New生成新的
 if x == nil && p.New != nil {
  x = p.New()
 }
 return x
}
```

## pin

```go
// pin pins the current goroutine to P, disables preemption and
// returns poolLocal pool for the P and the P's id.
// Caller must call runtime_procUnpin() when done with the pool.
func (p *Pool) pin() (*poolLocal, int) {
 pid := runtime_procPin()
 // In pinSlow we store to local and then to localSize, here we load in opposite order.
 // Since we've disabled preemption, GC cannot happen in between.
 // Thus here we must observe local at least as large localSize.
 // We can observe a newer/larger local, it is fine (we must observe its zero-initialized-ness).
 s := runtime_LoadAcquintptr(&p.localSize) // load-acquire
 l := p.local                              // load-consume
 if uintptr(pid) < s {
  return indexLocal(l, pid), pid
 }
 return p.pinSlow()
}
```

pin 方法里面首先会调用 runtime_procPin 方法会先获取当前 goroutine，然后绑定到对应的 M 上，然后返回 M 目前绑定的 P 的 id，因为这个 pid 后面会用到，防止在使用途中 P
被抢占，参考[golang 的对象池 sync.pool 源码解读](https://zhuanlan.zhihu.com/p/99710992)

之后使用原子操作取出 localSize，如果当前 pid 大于 localSize，那么就表示 Pool 还没创建对应的 poolLocal，那么调用 pinSlow 进行创建工作，否则调用 indexLocal 取出 pid 对应的
poolLocal 返回。

```go
func indexLocal(l unsafe.Pointer, i int) *poolLocal {
 lp := unsafe.Pointer(uintptr(l) + uintptr(i)*unsafe.Sizeof(poolLocal{}))
 return (*poolLocal)(lp)
}
```

indexLocal 里面是使用了地址操作，传入的 i 是数组的 index 值，所以需要获取 poolLocal{}的 size 做一下地址的位移操作，然后再转成转成 poolLocal 地址返回。

## pinSlow

```go
func (p *Pool) pinSlow() (*poolLocal, int) {
    // 解除pin
    runtime_procUnpin()
    // 加上全局锁
    allPoolsMu.Lock()
    defer allPoolsMu.Unlock()
    // pin住
    pid := runtime_procPin()
    s := p.localSize
    l := p.local
    // 重新对pid进行检查
    if uintptr(pid) < s {
        return indexLocal(l, pid), pid
    }
    // 初始化local前会将pool放入到allPools数组中
    if p.local == nil {
        allPools = append(allPools, p)
    }
    // 当前P的数量
    size := runtime.GOMAXPROCS(0)
    local := make([]poolLocal, size)
    atomic.StorePointer(&p.local, unsafe.Pointer(&local[0]))
    atomic.StoreUintptr(&p.localSize, uintptr(size))
    return &local[pid], pid
}
```

![pin](https://img.luozhiyun.com/20201226184352.png)

## popHead

```go
func (c *poolChain) popHead() (interface{}, bool) {
    // 这里头部是一个poolChainElt
    d := c.head
    // 遍历poolChain链表
    for d != nil {
        // 从poolChainElt的环状列表中获取值
        if val, ok := d.popHead(); ok {
            return val, ok
        }
        // load poolChain下一个对象
        d = loadPoolChainElt(&d.prev)
    }
    return nil, false
}
```

```go
func (d *poolDequeue) popHead() (interface{}, bool) {
    var slot *eface
    for {
        ptrs := atomic.LoadUint64(&d.headTail)
        // headTail的高32位为head，低32位为tail
        head, tail := d.unpack(ptrs)
        // 首尾相等，那么这个队列就是空的
        if tail == head {
            return nil, false
        }
        // 这里需要head--之后再获取slot
        head--
        ptrs2 := d.pack(head, tail)
        if atomic.CompareAndSwapUint64(&d.headTail, ptrs, ptrs2) {
            slot = &d.vals[head&uint32(len(d.vals)-1)]
            break
        }
    }
    val := *(*interface{})(unsafe.Pointer(slot))
    // 说明没取到缓存的对象，返回 nil
    if val == dequeueNil(nil) {
        val = nil
    }
    // 重置slot
    *slot = eface{}
    return val, true
}
```

- poolDequeue 的 popHead 方法首先会获取到 headTail 的值，然后调用 unpack 解包，headTail 是一个 64 位的值，高 32 位表示 head，低 32 位表示 tail。
- 判断 head 和 tail 是否相等，相等那么这个队列就是空的；
- 如果队列不是空的，那么将 head 减一之后再使用，因为 head 当前指的位置是空值，表示下一个新对象存放的位置；
- CAS 重新设值新的 headTail，成功之后获取 slot，这里因为 vals 大小是 2 的 n 次幂，因此`len(d.vals)-1)`之后低 n 位全是 1，和 head 取与之后可以获取到 head 的低 n 位的值；
- 如果 slot 所对应的对象是 dequeueNil，那么表示是空值，直接返回，否则将 slot 指针对应位置的值置空，返回 val。

## getSlow

```go
func (p *Pool) getSlow(pid int) interface{} {
    size := atomic.LoadUintptr(&p.localSize) // load-acquire
    locals := p.local                        // load-consume
    // 遍历locals列表，从其他的local的shared列表尾部获取对象
    for i := 0; i < int(size); i++ {
        l := indexLocal(locals, (pid+i+1)%int(size))
        if x, _ := l.shared.popTail(); x != nil {
            return x
        }
    }
    size = atomic.LoadUintptr(&p.victimSize)
    if uintptr(pid) >= size {
        return nil
    }
    locals = p.victim
    l := indexLocal(locals, pid)
    // victim的private不为空则返回
    if x := l.private; x != nil {
        l.private = nil
        return x
    }
    //  遍历victim对应的locals列表，从其他的local的shared列表尾部获取对象
    for i := 0; i < int(size); i++ {
        l := indexLocal(locals, (pid+i)%int(size))
        if x, _ := l.shared.popTail(); x != nil {
            return x
        }
    }
    // 获取不到，将victimSize置为0
    atomic.StoreUintptr(&p.victimSize, 0)
    return nil
}
```

## poolChain#popTail

```go
func (c *poolChain) popTail() (interface{}, bool) {
    d := loadPoolChainElt(&c.tail)
    // 如果最后一个节点是空的，那么直接返回
    if d == nil {
        return nil, false
    }

    for {
        // 这里获取的是next节点，与一般的双向链表是相反的
        d2 := loadPoolChainElt(&d.next)
        // 获取尾部对象
        if val, ok := d.popTail(); ok {
            return val, ok
        }

        if d2 == nil {
            return nil, false
        }
        // 因为d已经没有数据了，所以重置tail为d2，并删除d2的上一个节点
        if atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&c.tail)), unsafe.Pointer(d), unsafe.Pointer(d2)) {
            storePoolChainElt(&d2.prev, nil)
        }
        d = d2
    }
}
```

- 判断 poolChain，如果最后一个节点是空的，那么直接返回；
- 进入 for 循环，获取 tail 的 next 节点，这里需要注意的是这个双向链表与一般的链表是反向的；
- 调用 popTail 获取 poolDequeue 列表的对象，有对象直接返回；
- d2 为空则表示已经遍历完整个 poolChain 双向列表了，都为空，那么直接返回；
- 通过 CAS 将 tail 重置为 d2，因为 d 已经没有数据了，并将 d2 的 prev 节点置为 nil，然后将 d 置为 d2，进入下一个循环；

## poolDequeue#popTail

```go
func (d *poolDequeue) popTail() (interface{}, bool) {
    var slot *eface
    for {
        ptrs := atomic.LoadUint64(&d.headTail)
        // 和pophead一样，将headTail解包
        head, tail := d.unpack(ptrs)
        // 首位相等，表示列表中没有数据，返回
        if tail == head {
            return nil, false
        }
        ptrs2 := d.pack(head, tail+1)
        // CAS重置tail位置
        if atomic.CompareAndSwapUint64(&d.headTail, ptrs, ptrs2) {
            // 获取tail位置对象
            slot = &d.vals[tail&uint32(len(d.vals)-1)]
            break
        }
    }
    val := *(*interface{})(unsafe.Pointer(slot))
    // 判断对象是不是为空
    if val == dequeueNil(nil) {
        val = nil
    }
    // 将slot置空
    slot.val = nil
    atomic.StorePointer(&slot.typ, nil)
    return val, true
}
```

![pic](https://img.luozhiyun.com/20201226184358.png)

## Put

```go
// Put adds x to the pool.
func (p *Pool) Put(x interface{}) {
    if x == nil {
        return
    }
    ...
    l, _ := p.pin()
    if l.private == nil {
        l.private = x
        x = nil
    }
    if x != nil {
        l.shared.pushHead(x)
    }
    runtime_procUnpin()
    ...
}
func (c *poolChain) pushHead(val interface{}) {
    d := c.head
    // 头节点没有初始化，那么设值一下
    if d == nil {
        const initSize = 8 // Must be a power of 2
        d = new(poolChainElt)
        d.vals = make([]eface, initSize)
        c.head = d
        storePoolChainElt(&c.tail, d)
    }
    // 将对象加入到环状队列中
    if d.pushHead(val) {
        return
    }
    newSize := len(d.vals) * 2
    // 这里做了限制，单个环状队列不能超过2的30次方大小
    if newSize >= dequeueLimit {
        newSize = dequeueLimit
    }
    // 初始化新的环状列表，大小是d的两倍
    d2 := &poolChainElt{prev: d}
    d2.vals = make([]eface, newSize)
    c.head = d2
    storePoolChainElt(&d.next, d2)
    // push到新的队列中
    d2.pushHead(val)
}
func (d *poolDequeue) pushHead(val interface{}) bool {
    ptrs := atomic.LoadUint64(&d.headTail)
    // 解包headTail
    head, tail := d.unpack(ptrs)
    // 判断队列是否已满
    if (tail+uint32(len(d.vals)))&(1<<dequeueBits-1) == head {
        return false
    }
    // 找到head的槽位
    slot := &d.vals[head&uint32(len(d.vals)-1)]
    // 检查slot是否和popTail有冲突
    typ := atomic.LoadPointer(&slot.typ)
    if typ != nil {
        return false
    }
    if val == nil {
        val = dequeueNil(nil)
    }
    // 将 val 赋值到 slot，并将 head 指针值加 1
    *(*interface{})(unsafe.Pointer(slot)) = val
    atomic.AddUint64(&d.headTail, 1<<dequeueBits)
    return true
}
```

## GC

```go
func init() {
    runtime_registerPoolCleanup(poolCleanup)
}

func poolCleanup() {
    for _, p := range oldPools {
        p.victim = nil
        p.victimSize = 0
    }
    for _, p := range allPools {
        p.victim = p.local
        p.victimSize = p.localSize
        p.local = nil
        p.localSize = 0
    }
    oldPools, allPools = allPools, nil
}
```

poolCleanup 会在 STW 阶段被调用。主要是将 local 和 victim 作交换，那么不至于 GC 把所有的 Pool 都清空了，而是需要两个 GC 周期才会被释放。如果 sync.Pool
的获取、释放速度稳定，那么就不会有新的池对象进行分配。

package pool

import (
	"encoding/json"
	"sync"
)

type Student struct {
	Name   string
	Age    int8
	Remark [1024]byte
}

var stuPool = sync.Pool{New: func() interface{} { return new(Student) }}

var buf, _ = json.Marshal(Student{
	Name: "Abc",
	Age:  100,
})

func unmarsh() {
	stu := stuPool.Get().(*Student)
	_ = json.Unmarshal(buf, stu)
	stuPool.Put(stu)
}

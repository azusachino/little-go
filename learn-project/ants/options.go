package ants

import "time"

// 通过Option装饰Options对象 => Functional Options的应用
type Option func(opts *Options)

// 实际调用方法
func loadOptions(options ...Option) *Options {
	opts := new(Options)
	for _, o := range options {
		o(opts)
	}
	return opts
}

type Options struct {
	ExpiryDuration   time.Duration     // 定义了多久清理一次过期的协程
	PreAlloc         bool              // 定义了在初始化协程池时是否要先获取所有内存
	MaxBlockingTasks int               // 最大阻塞任务数量
	NonBlocking      bool              // 是否非阻塞
	PanicHandler     func(interface{}) // 异常处理器
	Logger           Logger            // 日志插件
}

// WithOptions accepts the whole options config.
func WithOptions(options Options) Option {
	return func(opts *Options) {
		*opts = options
	}
}

// WithExpiryDuration sets up the interval time of cleaning up goroutines.
func WithExpiryDuration(expiryDuration time.Duration) Option {
	return func(opts *Options) {
		opts.ExpiryDuration = expiryDuration
	}
}

// WithPreAlloc indicates whether it should malloc for workers.
func WithPreAlloc(preAlloc bool) Option {
	return func(opts *Options) {
		opts.PreAlloc = preAlloc
	}
}

// WithMaxBlockingTasks sets up the maximum number of goroutines that are blocked when it reaches the capacity of pool.
func WithMaxBlockingTasks(maxBlockingTasks int) Option {
	return func(opts *Options) {
		opts.MaxBlockingTasks = maxBlockingTasks
	}
}

// WithNonblocking indicates that pool will return nil when there is no available workers.
func WithNonblocking(nonblocking bool) Option {
	return func(opts *Options) {
		opts.NonBlocking = nonblocking
	}
}

// WithPanicHandler sets up panic handler.
func WithPanicHandler(panicHandler func(interface{})) Option {
	return func(opts *Options) {
		opts.PanicHandler = panicHandler
	}
}

// WithLogger sets up a customized logger.
func WithLogger(logger Logger) Option {
	return func(opts *Options) {
		opts.Logger = logger
	}
}

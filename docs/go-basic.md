# Basics

Go 放弃了传统的机遇操作系统线程的并发模型，采用了**用户层轻量级线程**，或者说是 coroutine。

## Go 的基本实现方式

一个 Go 程序对于操作系统来说只是一个用户层程序。操作系统的眼中只有线程，它甚至不知道 goroutine 的存在。goroutine 的调度全靠 Go 自己完成，实现 Go 程序内 goroutine 之间公平地竞争 CPU 资源的任务就落到了 Go 运行时头上。而将这些 goroutine 按照一定算法放到 CPU 上执行的程序就称为 goroutine 调度器（goroutine scheduler）。

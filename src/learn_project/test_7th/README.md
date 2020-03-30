# GO TEST

## testing (package)

- go test -bench . -cpuprofile cpu.out
- go tool pprof cpu.out (可以进入pprof交互窗口)

---

- 使用IDE查看代码覆盖
- 使用Go Test获取代码覆盖报告
- 使用go tool cover查看代码覆盖报告

---

- test.B的使用
- 使用pprof优化性能

---

- 表格驱动测试
- 代码覆盖率
- 性能优化工具

```go
var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "user error"},
	{errNotFound, 404, "Not Found"},
	{errNoPermission, 403, "Forbidden"},
	{errUnknown, 500, "Internal Server Error"},
	{noError, 200, "no error"},
}
```
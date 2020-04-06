# http

```go
package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)
func h1() {
	res, err := http.Get("http://www.baidu.com")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	s, e := httputil.DumpResponse(res, true)
	if e != nil {
		panic(e)
	}
	fmt.Println(s)
}
```

## basic

1. 使用http客户端发送请求
2. 使用http.Client控制请求头部等
3. 使用httputil简化工作

## http服务器的性能分析

- `import _ "net/http/pprof"`
- 访问/debug/pprof/
- 使用go tool pprof分析性能

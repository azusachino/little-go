package ch

// 抢占式
func main() {
	txtResult := make(chan string, 5)

	go func() { txtResult <- getTxt("res1.flysnow.org") }()
	go func() { txtResult <- getTxt("res2.flysnow.org") }()
	go func() { txtResult <- getTxt("res3.flysnow.org") }()
	go func() { txtResult <- getTxt("res4.flysnow.org") }()
	go func() { txtResult <- getTxt("res5.flysnow.org") }()

	println(<-txtResult)
}

func getTxt(host string) string {
	//省略网络访问逻辑，直接返回模拟结果
	//http.Get(host+"/1.txt")
	return host + "：模拟结果"
}

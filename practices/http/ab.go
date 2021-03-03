package main

import (
	"flag"
	"fmt"
	"math"
	"net/http"
	"sort"
	"sync"
	"time"
)

var url string
var m string

var c uint
var n uint

// init all flag values
func init() {
	flag.StringVar(&url, "url", "http://example.com", "地址")
	flag.StringVar(&m, "m", "GET", "请求方法")
	flag.UintVar(&c, "c", 10, "并发量")
	flag.UintVar(&n, "n", 10, "总请求量")
	flag.Parse()
}

func main() {
	ab := BenchWeb(url, m, n, c)
	fmt.Println("总请求数：", ab.Requests, " 成功数：", ab.CompleteRequests)
	fmt.Println("总耗时：", ab.TimeTotal)
}

// ApacheBench 性能测试工具
type ApacheBench struct {
	Url              string  // URL
	Concurrency      uint    // 并发量
	Requests         uint    // 总请求数
	TimeTotal        float64 // 总时长
	CompleteRequests uint    // 完成的请求数
	FailedRequests   uint    // 失败请求数
	ResponseTimeAvg  float64 // 平均响应时长
	ResponseTime95   float64 // 95分位响应时长
}

// 单次请求的结果
type onceReqReport struct {
	elapsed time.Duration
	success bool
}

// 发起单个http请求
func sendRequest(url string, method string, resultChan chan *onceReqReport, wg *sync.WaitGroup) {
	res := onceReqReport{}
	t0 := time.Now()
	request, _ := http.NewRequest(method, url, nil)
	client := http.Client{}
	resp, _ := client.Do(request)
	if resp.StatusCode >= 400 {
		res.success = false
	} else {
		res.success = true
	}
	t1 := time.Since(t0)
	res.elapsed = t1
	resultChan <- &res
	wg.Done()
}

// 发起批量请求，并等待所有结果返回
func concurRequest(concurrency uint, url string, method string, resultChan chan *onceReqReport) {
	wg := sync.WaitGroup{}
	for i := 0; i < int(concurrency); i++ {
		go sendRequest(url, method, resultChan, &wg)
		wg.Add(1)
	}
	wg.Wait()
}

// 收集数据
func collect(resChan chan *onceReqReport, result []*onceReqReport) {
	for item := range resChan {
		result = append(result, item)
	}
}

// 统计
func static(result []*onceReqReport) (failCount uint, time95, avg float64) {
	elapseSet := make([]float64, 0)
	failCount = uint(0)
	sum := 0.0 // 总耗时
	for _, r := range result {
		e := r.elapsed.Seconds()
		elapseSet = append(elapseSet, e)
		sum += e
		if !r.success {
			failCount++
		}
	}
	// 排序
	sort.Float64s(elapseSet)

	// 算95分位
	idx := math.Floor(float64(len(result)) * 0.95)
	time95 = elapseSet[int(idx)]
	// 算平均值
	avg = sum / float64(len(result))

	return
}

// web性能压测
func BenchWeb(url string, method string, requestNum uint, concurrency uint) *ApacheBench {
	// todo:检查url
	// 按并发数发起请求
	count := requestNum / concurrency // 需要发起并发的次数
	if requestNum%concurrency != 0 {  // 有余数就要多一次
		count++
	}

	resultChan := make(chan *onceReqReport, concurrency) // 结果通过Chan输送
	result := make([]*onceReqReport, 0)

	go collect(resultChan, result) // 收集结果

	t0 := time.Now()
	for i := 0; i < int(count); i++ {
		concurRequest(concurrency, url, method, resultChan) // 并发请求
	}
	tn := time.Since(t0)
	// 统计结果
	failCount, time95, avg := static(result)

	return &ApacheBench{
		Url:              url,
		Concurrency:      concurrency,
		Requests:         requestNum,
		TimeTotal:        tn.Seconds(),
		CompleteRequests: requestNum - failCount,
		FailedRequests:   failCount,
		ResponseTimeAvg:  avg,
		ResponseTime95:   time95,
	}
}

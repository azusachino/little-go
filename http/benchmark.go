package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {

	var d []map[string]interface{}
	d = append(d, map[string]interface{}{
		"app":  1000000001,
		"uid":  "ycpang",
		"rid":  "112233",
		"time": 1628577771123,
		"t":    2,
		"c":    2002,
		"v":    "100",
		"st":   0,
		"puid": "ythu",
	})
	c, _ := json.Marshal(d)
	start := time.Now()
	fmt.Println("grand beginning", start)
	for i := 0; i < 10000; i++ {
		s := time.Now()
		res, err := http.DefaultClient.Post("http://127.0.0.1:9530/api/v1/logging/collect/list/event", "application/json", bytes.NewBuffer(c))
		if err != nil {
			fmt.Printf("error : %v", err)
		}
		fmt.Println(time.Now(), res.Status, time.Since(s))

	}

	fmt.Println("Slow ending", time.Since(start))
	time.Sleep(1 * time.Second)
}

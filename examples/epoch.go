package examples

import (
	"fmt"
	"time"
)

func main() {

	// 2012-10-31 16:13:58.292387 +0000 UTC
	// 1351700038
	// 1351700038292
	// 1351700038292387000
	// 2012-10-31 16:13:58 +0000 UTC
	// 2012-10-31 16:13:58.292387 +0000 UTC
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}

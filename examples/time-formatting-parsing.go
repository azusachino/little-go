package examples

import (
	"fmt"
	"time"
)

// 2014-04-15T18:00:15-07:00
// 2012-11-01 22:08:41 +0000 +0000
// 6:00PM
// Tue Apr 15 18:00:15 2014
// 2014-04-15T18:00:15.161182-07:00
// 0000-01-01 20:41:00 +0000 UTC
// 2014-04-15T18:00:15-00:00
// parsing time "8:41PM" as "Mon Jan _2 15:04:05 2006": ...
func TimeF_() {
	var e error
	p := fmt.Println

	t := time.Now()
	_, _ = p(t.Format(time.RFC3339))

	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	_, _ = p(t1)

	if e != nil {
		panic(e)
	}
	_, _ = p(t.Format("3:04PM"))
	_, _ = p(t.Format("Mon Jan _2 15:04:05 2006"))
	_, _ = p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	_, _ = p(t2)
	if e != nil {
		panic(e)
	}

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	_, _ = p(e)

	if e != nil {
		panic(e)
	}
}

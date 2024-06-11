package examples

import (
	"fmt"
	"time"
)

func Time() {
	p := fmt.Println

	now := time.Now()
	_, _ = p(now)

	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	_, _ = p(then)

	_, _ = p(then.Year())
	_, _ = p(then.Month())
	_, _ = p(then.Day())
	_, _ = p(then.Hour())
	_, _ = p(then.Minute())
	_, _ = p(then.Second())
	_, _ = p(then.Nanosecond())
	_, _ = p(then.Location())

	_, _ = p(then.Weekday())

	_, _ = p(then.Before(now))
	_, _ = p(then.After(now))
	_, _ = p(then.Equal(now))

	diff := now.Sub(then)
	_, _ = p(diff)

	_, _ = p(diff.Hours())
	_, _ = p(diff.Minutes())
	_, _ = p(diff.Seconds())
	_, _ = p(diff.Nanoseconds())

	_, _ = p(then.Add(diff))
	_, _ = p(then.Add(-diff))
}

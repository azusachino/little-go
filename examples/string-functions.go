package examples

import (
	"fmt"
	"strings"
)

var p = fmt.Println

func init() {

	_, _ = p("Contains:  ", strings.Contains("test", "es"))
	_, _ = p("Count:     ", strings.Count("test", "t"))
	_, _ = p("HasPrefix: ", strings.HasPrefix("test", "te"))
	_, _ = p("HasSuffix: ", strings.HasSuffix("test", "st"))
	_, _ = p("Index:     ", strings.Index("test", "e"))
	_, _ = p("Join:      ", strings.Join([]string{"a", "b"}, "-"))
	_, _ = p("Repeat:    ", strings.Repeat("a", 5))
	_, _ = p("Replace:   ", strings.Replace("foo", "o", "0", -1))
	_, _ = p("Replace:   ", strings.Replace("foo", "o", "0", 1))
	_, _ = p("Split:     ", strings.Split("a-b-c-d-e", "-"))
	_, _ = p("ToLower:   ", strings.ToLower("TEST"))
	_, _ = p("ToUpper:   ", strings.ToUpper("test"))
	_, _ = p()

	_, _ = p("Len: ", len("hello"))
	_, _ = p("Char:", "hello"[1])
}

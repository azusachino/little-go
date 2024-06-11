package examples

import (
	"sync"
)

type Counter struct {
	sync.Mutex
	Count uint64
}

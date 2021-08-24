package pool

import (
	"encoding/json"
	"testing"
)

func BenchmarkUnmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := &Student{}
		_ = json.Unmarshal(buf, stu)
	}
}

func BenchmarkUnmarshalWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := stuPool.Get().(*Student)
		_ = json.Unmarshal(buf, stu)
		stuPool.Put(stu)
	}
}

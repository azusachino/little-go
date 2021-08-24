package main

import (
	"testing"
)

func benchmark(b *testing.B, f func(int, string) string) {
	var str = randomString(10)
	for i := 0; i < b.N; i++ {
		f(10000, str)
	}
}

//BenchmarkPlusConcat
//BenchmarkPlusConcat-8      	      14	  75699414 ns/op
//BenchmarkSprintfConcat
//BenchmarkSprintfConcat-8   	       8	 155750225 ns/op
//BenchmarkBuilderConcat
//BenchmarkBuilderConcat-8   	    7500	    150379 ns/op
//BenchmarkBufferConcat
//BenchmarkBufferConcat-8    	    7999	    139929 ns/op
//BenchmarkByteConcat
//BenchmarkByteConcat-8      	    7999	    147742 ns/op
//BenchmarkPreByteConcat
//BenchmarkPreByteConcat-8   	   14065	     71526 ns/op

func BenchmarkPlusConcat(b *testing.B)    { benchmark(b, plusConcat) }
func BenchmarkSprintfConcat(b *testing.B) { benchmark(b, sprintfConcat) }
func BenchmarkBuilderConcat(b *testing.B) { benchmark(b, builderConcat) }
func BenchmarkBufferConcat(b *testing.B)  { benchmark(b, bufferConcat) }
func BenchmarkByteConcat(b *testing.B)    { benchmark(b, byteConcat) }
func BenchmarkPreByteConcat(b *testing.B) { benchmark(b, preByteConcat) }

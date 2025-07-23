package main

import (
    "strings"
    "testing"
)

var testData = []string{"a", "b", "c", "d", "e", "f", "g"}

func BenchmarkStringPlus(b *testing.B) {
    b.ReportAllocs() // Reports memory allocations per operation
    for i := 0; i < b.N; i++ {
        var result string
        for _, s := range testData {
            result += s
        }
    }
}

func BenchmarkStringBuilder(b *testing.B) {
    b.ReportAllocs()
    for i := 0; i < b.N; i++ {
        var builder strings.Builder
        for _, s := range testData {
            builder.WriteString(s)
        }
        _ = builder.String()
    }
}

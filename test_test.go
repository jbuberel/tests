package main

import "testing"

var str = "Thompson"
var sub = "son"
func BenchmarkStr(b *testing.B) {
    for n := 0; n < b.N; n++ {
        substring(str, sub)
    }
}

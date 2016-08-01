package main

import (
	"log"
	"testing"

	//"github.com/go-stack/stack"
)

func BenchmarkCalle101(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//runtime.Caller(0)
		//_, f, l, k := runtime.Caller(0)
		//	fmt.Sprintf("%s %s", "Some", "test")
		log.Print("11")
	}
}

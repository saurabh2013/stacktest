package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
)

func main() {

	f, e := os.Create("cpu.prof")
	if e != nil {
		fmt.Println(e)
		return
	}

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	for i := 0; i < 10000; i++ {
		runtime.Caller(0)
		//fmt.Println(f, l, k)
	}
}

package main

import (
	"context"
	"flag"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

func TestFuncSubSub(loopCount int) {
	count := 0
	for i := 0; i < loopCount; i++ {
		count += i
	}
}

func TestFuncSub(loopCount int) {
	count := 0
	for i := 0; i < loopCount; i++ {
		for j := 0; j < loopCount; j++ {
			count += i*loopCount + j
		}
	}

	for i := 0; i < loopCount; i++ {
		TestFuncSubSub(loopCount)
	}
}

func TestFunc(loopCount int) {
	// labels := pprof.Labels("key", "value", "key2", "value2")

	count := 0
	for i := 0; i < loopCount; i++ {
		for j := 0; j < loopCount; j++ {
			for k := 0; k < loopCount; k++ {
				count += k
			}
		}
	}

	for i := 0; i < loopCount; i++ {
		TestFuncSub(loopCount)
	}
}

func TestFunc2(loopCount int) {
	count := 0
	for i := 0; i < loopCount; i++ {
		for j := 0; j < loopCount; j++ {
			for k := 0; k < loopCount; k++ {
				count += k
			}
		}
	}

	for i := 0; i < loopCount; i++ {
		TestFuncSub(loopCount)
	}
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {
	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	ctx := context.Background()
	pprof.Do(ctx, pprof.Labels("key", "value1"), func(ctx context.Context) {
		TestFunc(300)
	})
	pprof.Do(ctx, pprof.Labels("key", "value2"), func(ctx context.Context) {
		TestFunc2(600)
	})

	// ... rest of the program ...

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close()
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}

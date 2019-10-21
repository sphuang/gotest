package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func TestFunc2() {
	count := 0
	for i := 0; i < 500; i++ {
		for j := 0; j < 500; j++ {
			count += (i - j)
		}
	}
}

func TestFunc1() {
	count := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			count += (i - j)
		}
	}
}

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	for {
		time.Sleep(time.Second)
		TestFunc1()
		TestFunc2()
	}
}

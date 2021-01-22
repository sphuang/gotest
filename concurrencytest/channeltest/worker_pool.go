package channeltest

import (
	"fmt"
	"time"
	"runtime"
	"bytes"
	"strconv"
	"math/rand"
)

type WorkerManager struct {
	NumWorkers int
}

func (m *WorkerManager) Start() {
	c := make(chan bool, m.NumWorkers)

	for i := 0; i < m.NumWorkers; i++ {
		m.launchWorker(c)
	}

	for {
		select {
		case <- c:
			m.launchWorker(c)
		}
	}
}

func (m *WorkerManager) launchWorker(c chan<- bool) {
	w := &Worker{}
	go w.Run(c)
}

type Worker struct {}

func (w *Worker) Run(c chan<- bool) {
	fmt.Printf("Routine [%v] has started!\n", getGID())
	time.Sleep(time.Second * time.Duration(rand.Float64() * 30))
	fmt.Printf("Routine [%v] has finished!\n", getGID())
	c <- true
}

func getGID() uint64 {
    b := make([]byte, 64)
    b = b[:runtime.Stack(b, false)]
    b = bytes.TrimPrefix(b, []byte("goroutine "))
    b = b[:bytes.IndexByte(b, ' ')]
    n, _ := strconv.ParseUint(string(b), 10, 64)
    return n
}
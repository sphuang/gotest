package contexttest

import (
	"fmt"
	"sync"
	"time"
	"bytes"
	"strconv"
	"runtime"
	"context"
)

func subtask(ctx context.Context, wg *sync.WaitGroup) {
	fmt.Printf("Subtask [%v] Start\n", getGID())
	for {
		select {
		case <- ctx.Done():
			fmt.Printf("Subtask [%v] End\n", getGID())
			wg.Done()
			return
		}
	}
}

func mainloop(wg *sync.WaitGroup) {
	boom := time.After(5 * time.Second)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := 0; i < 5; i++ {
		go subtask(ctx, wg)
		wg.Add(1)
	}

	for {
		select {
		case <-boom:
			fmt.Println("Main loop: BOOM!")
			return
		}
	}
}

func GracefulShutdown() {
	var wg sync.WaitGroup

	mainloop(&wg)

	wg.Wait()
	fmt.Println("Terminate after all subtasks are shutdown.")
}

func getGID() uint64 {
    b := make([]byte, 64)
    b = b[:runtime.Stack(b, false)]
    b = bytes.TrimPrefix(b, []byte("goroutine "))
    b = b[:bytes.IndexByte(b, ' ')]
    n, _ := strconv.ParseUint(string(b), 10, 64)
    return n
}
package channeltest

import (
	"testing"
)

func TestWorkerPool(t *testing.T) {
	m := WorkerManager{3}
	m.Start()
}
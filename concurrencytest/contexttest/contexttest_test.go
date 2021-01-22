package contexttest

import (
	// "log"
	// "reflect"
	"time"
	"testing"
	// "github.com/stretchr/testify/assert"
)

func TestAPITimeout(t *testing.T) {
	// will timeout
	LaunchAPI(3 * time.Second)

	// will finish
	LaunchAPI(6 * time.Second)
}

func TestGracefulShutdown(t *testing.T) {
	GracefulShutdown()
}
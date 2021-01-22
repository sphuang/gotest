package channeltest

import (
	"testing"
	// "github.com/stretchr/testify/assert"
)

func TestIOThreadSendMessage(t *testing.T) {
	SendMessage("Hello")
}

func TestIOThreadSendMessageNoWait(t *testing.T) {
	SendMessageNoWait("Hello")
}
package channeltest

import (
	"time"
	"fmt"
)

func doSendMessage(msg string, c chan<- bool) {
	time.Sleep(time.Second)
	fmt.Printf("Message %v has Sended", msg)
	c <- true
}

func SendMessage(msg string) {
	c := make(chan bool)
	go doSendMessage(msg, c)
	<-c
	fmt.Println("Done")
}

func SendMessageNoWait(msg string) {
	go func(msg string) {
		time.Sleep(time.Second)
		fmt.Printf("Message %v has Sended", msg)
	} (msg)
	fmt.Println("Done")
}
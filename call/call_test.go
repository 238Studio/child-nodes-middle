package call_test

import (
	"log"
	"testing"

	"github.com/238Studio/child-nodes-middle/call"
)

func TestCall(t *testing.T) {
	call.InitManager()
	go send()
	get()
}

func send() {
	_ = call.InitReceiver("test")
	re := call.Call("test1", "test", 42)
	log.Println(re)
}

func get() {
	receiver := call.InitReceiver("test1")
	for {
		select {
		case callForm := <-receiver.GetCall:
			log.Println(callForm.GetMethodName())
			callForm.Return(callForm.GetArgs())
		}
	}
}

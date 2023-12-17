package call_test

import (
	"log"
	"testing"
	"unsafe"

	"github.com/238Studio/child-nodes-middle/call"
)

func TestCall(t *testing.T) {
	call.InitManager()
	var stop = make(chan struct{})
	go send(stop)
	get(stop)
}

func send(stop chan struct{}) {
	_ = call.InitReceiver("test")
	arg := make([]uintptr, 1)
	a := "42"
	arg[0] = uintptr(unsafe.Pointer(&a))
	re := call.Call("test1", "test", arg)
	log.Println(*(*string)(unsafe.Pointer(re)))
	stop <- struct{}{}
	log.Println("测试完成")
}

func get(stop chan struct{}) {
	receiver := call.InitReceiver("test1")
	for {
		select {
		case callForm := <-receiver.GetCall:
			log.Println(callForm.GetMethodName())
			callForm.Return(callForm.GetArgs()[0])
		case <-stop:
			return
		}
	}
}

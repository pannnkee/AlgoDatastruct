package Semaphore

import (
	"fmt"
	"sync/atomic"
	"time"
)

type NonBlockSemaphore struct {
	interCapacity int64
}

var capacity = 5

func NewNonBlockSemaphore() *NonBlockSemaphore {
	return &NonBlockSemaphore{
		interCapacity: int64(capacity),
	}
}

func (this *NonBlockSemaphore) Acquire() {
	if atomic.LoadInt64(&this.interCapacity) > 0 {
		atomic.AddInt64(&this.interCapacity, -1)
		fmt.Println("get Semaphore")
		//do something
		return
	} else {
		fmt.Println("semaphore acquire is blocking")
	}
}

func (this *NonBlockSemaphore) Release() {
	if atomic.LoadInt64(&this.interCapacity) < int64(capacity) {
		atomic.AddInt64(&this.interCapacity, 1)
	}

}

func (this *NonBlockSemaphore) ShowCapacity() {
	if this.interCapacity > int64(capacity) {
		fmt.Println(this.interCapacity)
	}
}

func NonBlockingInstance() {
	semaphore := NewNonBlockSemaphore()
	go func() {
		for {
			semaphore.Acquire()
			time.Sleep(time.Millisecond * 500)
		}
	}()
	go func() {
		for {
			semaphore.Release()
			time.Sleep(time.Second)
		}
	}()
	go func() {
		for {
			semaphore.ShowCapacity()
		}
	}()
	time.Sleep(time.Hour)
}
package Semaphore

import (
	"fmt"
	"time"
)

type BlockSemaphore struct {
	innerChan chan struct{}
}

func NewSemaphore(capacity int64) *BlockSemaphore {
	return &BlockSemaphore{
		innerChan: make(chan struct{}, capacity),
	}
}

func (this *BlockSemaphore) Acquire() {
	select {
		case this.innerChan <- struct{}{}:
			fmt.Println("acquire  semaphore")
			return
		default:
			fmt.Println("semaphore acquire is blocking")
			time.Sleep(time.Millisecond * 100)
	}
}

func (this *BlockSemaphore) Release() {
	select {
	case <-this.innerChan :
		fmt.Println("release")
		return
	default:
		return
	}
}

func Instance() {
	semaphore := NewSemaphore(5)
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
	time.Sleep(time.Hour)
}



package stack

import (
	"fmt"
	"sync"
)

type Item interface {}

type SliceStack struct {
	items []Item
	lock sync.RWMutex
}

func NewStackSlice() *SliceStack {
	return &SliceStack{items: []Item{}}
}

func (this *SliceStack) Pop() Item {
	this.lock.Lock()
	defer this.lock.Unlock()

	if len(this.items) == 0 {
		fmt.Print("stack is nil, return.")
		return nil
	}

	Item := this.items[len(this.items) - 1]
	this.items = this.items[0 : len(this.items) - 1]
	return Item
}

func (this *SliceStack) Push(t Item) {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.items = append(this.items, t)
}

func (this *SliceStack) Print() {
	fmt.Println(this.items)
}

func ItemStackInterface() {
	stack := NewStackSlice()
	for i := 0; i < 5; i++ {
		stack.Push(i)
	}
	stack.Print()

	for i := 0; i < 3; i++ {
		stack.Pop()
	}
	stack.Print()
}


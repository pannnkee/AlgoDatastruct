package src

import (
	"fmt"
	"sync"
)

type Item interface {}

type ItemStack struct {
	items []Item
	lock sync.RWMutex
}

//新建一个栈
func newItemStack() *ItemStack {
	s := &ItemStack{
		items: []Item{},
	}
	return s
}

//打印
func (this *ItemStack) Print() {
	fmt.Println(this.items)
}

//栈添加元素
func (this *ItemStack) Push(t Item) {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.items = append(this.items, t)
}

//输出元素
func (this *ItemStack) Pop() Item {
	this.lock.Lock()
	defer this.lock.Unlock()

	if len(this.items) == 0 {
		return nil
	}
	item := this.items[len(this.items) - 1]
	this.items = this.items[0:len(this.items) - 1]
	return item
}

func ItemStackInstance() {
	stack := newItemStack()
	for i := 0; i < 10; i++ {
		stack.Push(i)
	}
	stack.Print()

	for i := 0; i < 3; i++ {
		stack.Pop()
	}
	stack.Print()
}



package stack

import (
	"container/list"
	"fmt"
)

type ListStack struct {
	list *list.List
}

func NewListStack() *ListStack {
	list := list.New()
	return &ListStack{list: list}
}

func (this *ListStack) Push(value interface{}) {
	this.list.PushBack(value)
}

func (this *ListStack) Pop() interface{} {
	back := this.list.Back()
	if back != nil {
		this.list.Remove(back)
		return back.Value
	}
	return nil
}

func (this *ListStack) Print() {
	for e := this.list.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
	}
	fmt.Println()
}

func ListStackInstance() {
	stackList := NewListStack()
	for i := 0; i < 5; i++ {
		stackList.Push(i)
	}
	stackList.Print()

	for i := 0; i < 3; i++ {
		stackList.Pop()
	}
	stackList.Print()
}
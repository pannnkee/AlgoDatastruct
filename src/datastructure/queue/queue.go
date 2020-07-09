package queue

import (
	"container/list"
	"fmt"
)

type Queue struct {
	List *list.List
}

func (this *Queue) Push(value interface{}) {
	this.List.PushBack(value)
}

func (this *Queue) Pop() interface{} {
	item := this.List.Front()
	if item != nil {
		this.List.Remove(item)
		return item
	}
	return nil
}

func (this *Queue) Print() {
	for e := this.List.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
	}
	fmt.Println()
}

func NewQueue() *Queue {
	return &Queue{ List: list.New() }
}



func InstanceQueue() {
	q := NewQueue()
	for i := 0; i < 5; i++ {
		q.Push(i)
	}
	q.Print()
	for i := 0; i < 2; i++ {
		q.Pop()
	}
	q.Print()
}

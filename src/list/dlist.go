package list

import (
	"fmt"
)

type DNode struct {
	Data Element
	Prev *DNode
	Next *DNode
}

type DList struct {
	Head *DNode
	Tail *DNode
	Size  int
}

//初始化链表
func newDList() *DList {
	return &DList{Head: nil, Tail: nil, Size: 0}
}

func (this *DList) Length() int  {
	return this.Size
}

func newDListNode(data Element) *DNode {
	return &DNode{Data: data}
}

//头插法
func (this *DList) DListHeadAdd(data Element) {

	node := newDListNode(data)
	if this.Length() == 0 {
		this.Head = node
		this.Tail = node
		node.Prev = nil
		node.Next = nil
	} else {
		node.Next = this.Head
		this.Head.Prev = node
		this.Head = node
	}
	this.Size++
}

//尾插法
func (this *DList) DListTailAdd(data Element) {
	node := newDListNode(data)
	if this.Length() == 0 {
		this.Tail = node
		this.Head = node
		node.Prev = nil
		node.Next = nil
	} else {
		node.Prev = this.Tail
		node.Next = nil
		this.Tail.Next = node
		this.Tail = node
	}
	this.Size++
}

//任意位置插入
func (this *DList) InsertAnyWhere(index int, data Element) {
	if index > this.Length() {
		fmt.Println("Index error, Please check index again.")
		return
	}
	node := newDListNode(data)
	if this.Length() == 0 {
		this.Head = node
		this.Tail = node
		node.Prev = nil
		node.Next = nil
		this.Size++
		return
	}
	if index == 0 {
		this.DListHeadAdd(data)
		return
	}
	if index == this.Length() - 1 {
		this.DListTailAdd(data)
		return
	}

	head := this.Head
	for i := 0; i < index - 1; i++ {
		head = head.Next
	}
	node.Next = head.Next
	head.Next = node

	node.Prev = head
	head.Next.Prev = node
	this.Size++
}

//任意位置删除
func (this *DList) DeleteAnywhere(index int) {
	if index > this.Size {
		fmt.Println("index error, please check again.")
		return
	}

	p := this.Head
	for i := 0; i < index - 1; i++ {
		p = p.Next
	}
	p.Next = p.Next.Next
	p.Next.Next.Prev = p
	this.Size--
}
func (this *DList) PrintDList() {
	head := this.Head
	for head != nil {
		fmt.Printf(fmt.Sprintf(" %d ", head.Data))
		head = head.Next
	}
	fmt.Println()
}

func DListNodeInstance() {
	DList := newDList()
	for i := 0; i < 5; i++ {
		//DList.DListHeadAdd(Element(i))
		DList.DListTailAdd(Element(i))
	}
	DList.PrintDList()
	fmt.Println("size:", DList.Size)

	DList.InsertAnyWhere(1,100)
	DList.PrintDList()

	DList.InsertAnyWhere(2,500)
	DList.PrintDList()

	DList.InsertAnyWhere(4,600)
	DList.PrintDList()

	DList.DeleteAnywhere(3)
	DList.PrintDList()

	DList.DeleteAnywhere(3)
	DList.PrintDList()
}

package src

import "fmt"

type DNode struct {
	Data Element
	Prev *DNode
	Next *DNode
}

type DList struct {
	Head *DNode
	Tail *DNode
}

func newDList() *DList {
	return &DList{Head: nil, Tail: nil}
}

func (this *DList) Length() (count int) {
	head := this.Head
	if head != nil {
		count++
		head = head.Next
	}
	return
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
	} else {
		node.Next = this.Head
		this.Head.Prev = node
		this.Head = node
	}
}

//尾插法
func (this *DList) DListTailAdd(data Element) {
	node := newDListNode(data)
	if this.Length() == 0 {
		this.Tail = node
		this.Head = node
	} else {

		node.Prev = this.Tail
		this.Tail.Next = node
		this.Tail = node
	}
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

	DList.InsertAnyWhere(2,100)
	DList.PrintDList()

	DList.InsertAnyWhere(0,50)
	DList.PrintDList()
	DList.InsertAnyWhere(3,60)
	DList.PrintDList()
}

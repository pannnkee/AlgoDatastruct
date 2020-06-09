package src

import "fmt"

type Element int64

type ListNode struct {
	Data Element
	Next *ListNode
}

func newListNode(data Element) *ListNode {
	return &ListNode{Data: data}
}

func newHeader() *ListNode {
	return &ListNode{Data: 0, Next: nil,}
}

//从尾部插入
func TailAdd(head *ListNode, data Element) {
	point := head
	for point.Next != nil {
		point = point.Next
	}

	node := newListNode(data)
	point.Next = node
}

//从头部插入
func HeadAdd(head *ListNode, data Element) {
	point := head

	node := newListNode(data)
	node.Next = point.Next
	point.Next = node
}

//指定位置插入
func Insert(head *ListNode, index int, data Element) {
	if index < 0 || index > GetListLength(head) {
		fmt.Println("Index error,Please check again.")
		return
	}
	point := head
	for i := 0; i < index - 1 ; i++ {
		point = point.Next
	}

	node := newListNode(data)
	node.Next = point.Next
	point.Next = node
}

//删除指定Index元素
func Delete(head *ListNode, index int) {
	if index < 0 ||  index > GetListLength(head) {
		fmt.Println("Index error,Please check again.")
		return
	}

	point := head
	for i := 0; i < index - 1; i++ {
		point = point.Next
	}
	point.Next = point.Next.Next
}

// 查询元素的位置
func Search(head *ListNode, data Element) (index int) {
	point := head
	index = 0
	for point.Next != nil {
		if point.Data == data {
			fmt.Println(fmt.Sprintf("Search data Done, Data:%d Index:%d", data, index))
			return
		} else {
			index++
			point = point.Next
			if index > GetListLength(head) - 1 {
				fmt.Println(fmt.Sprintf("Data %d not exist", data))
				break
			}
		}
	}
	return
}

//根据Index获取元素的值
func GetData(head *ListNode, index int) (data Element) {
	if index < 0 ||  index > GetListLength(head) {
		fmt.Println("Index error,Please check again.")
		return
	}

	point := head
	for i := 0; i < index; i++ {
		point = point.Next
	}
	data = point.Data
	fmt.Println(fmt.Sprintf("Index:%d Data:%d", index, data))
	return
}

func (head *ListNode) PrintList() {
	point := head
	if point == nil {
		fmt.Println("List is nil, return!")
		return
	}
	for point.Next != nil {
		fmt.Print(fmt.Sprintf(" %d ",point.Data))
		point = point.Next
	}
	fmt.Println()
}

func GetListLength(head *ListNode) (length int) {
	point := head
	for point.Next != nil {
		length++
		point = point.Next
	}
	return
}

func ListInstance()  {
	Head := newHeader()
	if Head == nil {
		fmt.Println("Head is nil, return.")
	}

	for i := 1; i < 10; i++ {
		TailAdd(Head, Element(i))
	}
	Head.PrintList()

	HeadAdd(Head,1000)
	HeadAdd(Head,2000)
	HeadAdd(Head,3000)
	Head.PrintList()

	Delete(Head,2)
	Head.PrintList()

	Insert(Head,2, 101)
	Head.PrintList()

	Search(Head,101)
	GetData(Head,4)

}




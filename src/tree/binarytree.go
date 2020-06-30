package tree

import "fmt"

type Item interface {}

type BinaryTreeNode struct {
	//Key int
	Value Item
	Left *BinaryTreeNode
	Right *BinaryTreeNode
}

func (this *BinaryTreeNode) Print() {
	fmt.Print(this.Value, " ")
}

func (this *BinaryTreeNode) SetValue(value Item) {
	if this == nil {
		fmt.Println("setting value to nil.node ignored.")
		return
	}
	this.Value = value
}

// 前序遍历
func (this *BinaryTreeNode) PreOrder() {
	if this == nil {
		return
	}
	this.Print()
	this.Left.PreOrder()
	this.Right.PreOrder()
}

// 中序遍历
func (this *BinaryTreeNode) MiddleOrder() {
	if this == nil {
		fmt.Println("Tree is nil, return.")
		return
	}
	this.Left.MiddleOrder()
	this.Print()
	this.Right.MiddleOrder()
}

// 后续遍历
func (this *BinaryTreeNode) PostOrder() {
	if this == nil {
		fmt.Println("Tree is nil, return.")
		return
	}
	this.Left.PostOrder()
	this.Right.PostOrder()
	this.Print()
}

// 广度优先遍历
func (this *BinaryTreeNode) BreadthFirstSearch() {
	if this == nil {
		fmt.Println("Tree is nil, return.")
		return
	}
	result := []Item{}
	BinaryTreeNodes := []*BinaryTreeNode{this}
	for len(BinaryTreeNodes) > 0 {
		curBinaryTreeNode := BinaryTreeNodes[0]
		BinaryTreeNodes = BinaryTreeNodes[1:]
		result = append(result, curBinaryTreeNode.Value)
		if curBinaryTreeNode.Left != nil {
			BinaryTreeNodes = append(BinaryTreeNodes, curBinaryTreeNode.Left)
		}
		if curBinaryTreeNode.Right != nil {
			BinaryTreeNodes = append(BinaryTreeNodes, curBinaryTreeNode.Right)
		}
	}

	for _, v := range result {
		fmt.Print(v, " ")
	}
}

//层数
func (this *BinaryTreeNode) Layers() int {
	if this == nil {
		return 0
	}
	leftLayers := this.Left.Layers()
	rightLayers := this.Right.Layers()
	if leftLayers > rightLayers {
		return  leftLayers + 1
	} else {
		return  rightLayers + 1
	}
}

func CreateNode(v Item) *BinaryTreeNode {
	return &BinaryTreeNode{Value: v}
}

func BinaryTreeNodeInstance() {
	root := BinaryTreeNode{Value: 3}

	root.Left = &BinaryTreeNode{}
	root.Left.SetValue(0)
	root.Left.Right = CreateNode(2)

	root.Right = &BinaryTreeNode{5, nil, nil}
	root.Right.Left = CreateNode(4)

	fmt.Print("\n前序遍历:")
	root.PreOrder()

	fmt.Print("\n层数:", root.Layers())
}



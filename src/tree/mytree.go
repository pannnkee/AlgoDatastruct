package tree

import (
	"fmt"
	"sync"
)

type MyTreeNode struct {
	Key int
	Value Item
	Left *MyTreeNode
	Right *MyTreeNode
}

type MyTree struct {
	Lock sync.Mutex
	Root *MyTreeNode
}

func (this *MyTree) Print() {
	fmt.Println(this.Root.Value)
}

//插入节点
func (this *MyTree) Insert(key int, value Item) {
	this.Lock.Lock()
	defer this.Lock.Unlock()

	newNode := &MyTreeNode{Key: key, Value: value}
	if this.Root == nil {
		this.Root = newNode
	} else {
		Insert(this.Root, newNode)
	}
}

func Insert(node, newNode *MyTreeNode) {
	if newNode.Key < node.Key {
		if node.Left == nil {
			node.Left = node
		} else {
			Insert(node.Left, newNode)
		}
	} else {
		if node.Right == nil {
			node.Right = node
		} else {
			Insert(node.Right, newNode)
		}
	}
}

// 搜索节点
func (this *MyTree) Search(key int) bool {
	this.Lock.Lock()
	defer this.Lock.Unlock()
	return Search(this.Root, key)
}

func Search(node *MyTreeNode, key int) bool {
	if node == nil {
		return false
	}

	if key < node.Key {
		return Search(node.Left, key)
	}

	if key > node.Key {
		return Search(node.Right, key)
	}
	return true
}

// 删除节点
func (this *MyTree) Remove(key int) {
	this.Lock.Lock()
	defer this.Lock.Unlock()
	Remove(this.Root, key)
}

func Remove(node *MyTreeNode, key int) *MyTreeNode {
	if node == nil {
		return nil
	}

	//寻找节点
	//要删除的在左边
	if key < node.Key {
		node.Left = Remove(node.Left, key)
		return node
	}

	//要删除的在右边
	if key > node.Key {
		node.Right = Remove(node.Right, key)
		return node
	}

	//没有叶子节点
	if node.Left == nil && node.Right == nil {
		node = nil
		return node
	}

	//要删除的节点只有一个子节点 删除自身
	if node.Left == nil {
		node = node.Right
		return node
	}
	if node.Right == nil {
		node = node.Left
		return node
	}

	mostLeftNode := node.Right
	for  {
		if mostLeftNode != nil && mostLeftNode.Left != nil {
			mostLeftNode = mostLeftNode.Left
		} else {
			break
		}
	}
	// 使用右子树的最左节点替换当前节点，即删除当前节点
	node.Key, node.Value = mostLeftNode.Key, mostLeftNode.Value
	node.Right = Remove(node.Right, node.Key)
	return  node
}

// 最小节点
func (this *MyTree) Min() *Item {
	this.Lock.Lock()
	defer this.Lock.Unlock()

	node := this.Root
	if node == nil {
		return nil
	}
	for  {
		if node.Left == nil {
			return &node.Value
		}
		node = node.Left
	}
}

// 最大节点
func (this *MyTree) Max() *Item {
	this.Lock.Lock()
	defer this.Lock.Unlock()

	node := this.Root
	if node == nil {
		return nil
	}

	for  {
		if node.Right == nil {
			return &node.Value
		}
		node = node.Right
	}
}



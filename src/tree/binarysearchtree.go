package tree

import (
	"DataStructureGolang/pkg"
	"fmt"
	"sync"
)

type BinarySearchTreeNode struct {
	Key int
	Value Item
	Left *BinarySearchTreeNode
	Right *BinarySearchTreeNode
}

type BinarySearchTree struct {
	root *BinarySearchTreeNode
	lock sync.RWMutex
}

func (this *BinarySearchTree) Print() {
	fmt.Println(this.root.Value)
}

func (this *BinarySearchTree) Insert(key int, value Item) {
	this.lock.Lock()
	defer this.lock.Unlock()

	newNode := &BinarySearchTreeNode{Key:   key, Value: value}
	if this.root == nil {
		this.root = newNode
	} else {
		// 在树中递归查找正确的位置并插入
		InsertNode(this.root, newNode)
	}
}

func InsertNode(node, newNode *BinarySearchTreeNode) {
	//插入到左子树
	if newNode.Key < node.Key {
		if node.Left == nil {
			node.Left = newNode
		} else {
			//递归查找左边插入
			InsertNode(node.Left, newNode)
		}
	} else {
		if node.Right == nil {
			node.Right = newNode
		} else {
			InsertNode(node.Right, newNode)
		}
	}
}

// 搜索序号
func (this *BinarySearchTree) Search(key int) bool {
	this.lock.Lock()
	defer this.lock.Unlock()
	return  search(this.root, key)
}

func search(node *BinarySearchTreeNode, key int) bool {
	if node == nil {
		return false
	}

	if key < node.Key {
		return search(node.Left, key)
	}

	if key > node.Key {
		return search(node.Right, key)
	}
	return true
}

//删除节点
func (this *BinarySearchTree) Remove(key int) {
	this.lock.Lock()
	defer this.lock.Unlock()
	remove(this.root, key)
}

func remove(node *BinarySearchTreeNode, key int) *BinarySearchTreeNode {
	if node == nil {
		return nil
	}

	//寻找节点
	//要删除的节点在左侧
	if key < node.Key {
		node.Left = remove(node.Left, key)
		return node
	}

	//要删除的节点在右侧
	if key > node.Key {
		node.Right = remove(node.Right, key)
		return  node
	}

	//判断节点类型
	//要删除的节点是叶子节点， 直接删除
	if node.Left == nil && node.Right == nil {
		node = nil
		return node
	}

	//要删除的节点只有一个节点，删除自身
	if node.Left == nil {
		node = node.Right
		return node
	}
	if node.Right == nil {
		node = node.Left
		return node
	}

	// 要删除的节点有2个子节点 找到右子树的最左节点，替换当前节点
	mostLeftNode := node.Right
	for {
		// 一直遍历找到最左节点
		if mostLeftNode != nil && mostLeftNode.Left != nil {
			mostLeftNode = mostLeftNode.Left
		} else {
			break
		}
	}
	// 使用右子树的最左节点替换当前节点，即删除当前节点
	node.Key, node.Value = mostLeftNode.Key, mostLeftNode.Value
	node.Right = remove(node.Right, node.Key)
	return  node
}

// 获取树中值最小的节点: 最左节点
func (this *BinarySearchTree) Min() *Item {
	this.lock.Lock()
	defer this.lock.Unlock()

	node := this.root
	if node == nil {
		return  nil
	}
	for  {
		if node.Left == nil {
			return &node.Value
		}
		node = node.Left
	}
}

// 获取树中值最大的节点：最右节点
func (this *BinarySearchTree) Max() *Item {
	this.lock.Lock()
	defer this.lock.Unlock()

	node := this.root
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

//先序遍历: 根左右
func (this *BinarySearchTree) PreOrderTraverse(printFunc func(Item)) {
	this.lock.Lock()
	defer this.lock.Unlock()
	preOrderTraverse(this.root, printFunc)
}

func preOrderTraverse(node *BinarySearchTreeNode, printFunc func(Item)) {
 	if node !=  nil {
		printFunc(node.Value)
		preOrderTraverse(node.Left, printFunc)
		preOrderTraverse(node.Right, printFunc)
	}
}

//中序遍历: 左根右
func (this *BinarySearchTree) PostOrderTraverse(printFunc func(Item)) {
	this.lock.Lock()
	defer this.lock.Unlock()
	postOrderTraverse(this.root, printFunc)
}

func postOrderTraverse(node *BinarySearchTreeNode, printFunc func(Item)) {
	if node != nil {
		postOrderTraverse(node.Left, printFunc)
		postOrderTraverse(node.Right, printFunc)
		printFunc(node.Value)
	}
}

//后续遍历: 左右根
func (this *BinarySearchTree) InOrderTraverse(printFunc func(Item)) {
	this.lock.Lock()
	defer this.lock.Unlock()
	inOrderTraverse(this.root, printFunc)
}

func inOrderTraverse(node *BinarySearchTreeNode, printFunc func(Item)) {
	if node != nil {
		inOrderTraverse(node.Left, printFunc)
		printFunc(node.Value)
		inOrderTraverse(node.Right, printFunc)
	}
}

// 打印树结构
func (this *BinarySearchTree) String() {
	this.lock.Lock()
	defer this.lock.Unlock()
	if this.root == nil {
		println("Tree is empty")
		return
	}
	stringify(this.root, 0)
	println("----------------------------")
}
func stringify(node *BinarySearchTreeNode, level int) {
	if node == nil {
		return
	}

	format := ""
	for i := 0; i < level; i++ {
		format += "\t" // 根据节点的深度决定缩进长度
	}
	format += "----[ "
	level++
	// 先递归打印左子树
	stringify(node.Left, level)
	fmt.Printf(format+"%d\n", node.Key)
	/// 再递归打印右子树
	stringify(node.Right, level)
}

func BinarySearchTreeInstance() {
	searchTree := BinarySearchTree{}
	//searchTree.Insert(8,"8")

	number := pkg.GenerateRandomNumber(0, 100, 20)
	for i := 0; i < len(number); i++ {
		searchTree.Insert(number[i], number[i])
	}

	searchTree.PreOrderTraverse(func(item Item) {
		fmt.Print(" ", item)
	})

	fmt.Println()

	searchTree.Insert(11,"11")
	searchTree.PreOrderTraverse(func(item Item) {
		fmt.Print(" ", item)
	})

	fmt.Println()

	searchTree.Remove(4)
	searchTree.PreOrderTraverse(func(item Item) {
		fmt.Print(" ", item)
	})

	fmt.Println()

	searchTree.String()
}





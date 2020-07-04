package tree

import (
	"DataStructureGolang/pkg"
	"testing"
)

var bst BinarySearchTree

func fillTree(bst *BinarySearchTree) {
	number := pkg.GenerateRandomNumber(0, 100, 20)
	for i := 0; i < len(number); i++ {
		bst.Insert(number[i], number[i])
	}
}

func BenchmarkBinarySearchTree_Insert(b *testing.B) {
	number := pkg.GenerateRandomNumber(0, b.N, b.N/2)
	for i := 0; i < len(number); i++ {
		bst.Insert(number[i], number[i])
	}
}

func TestBinarySearchTree_Insert(t *testing.T) {
	fillTree(&bst)
	bst.String()

	bst.Insert(101,"101")
	bst.String()
}

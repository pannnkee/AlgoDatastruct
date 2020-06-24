package stack

import "testing"

var listStack *ListStack

func init() {
	listStack = NewListStack()
}

func BenchmarkListStack_Pop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		listStack.Push("test")
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		listStack.Pop()
	}
}

func BenchmarkListStack_Push(b *testing.B) {
	for i := 0; i < b.N ; i++ {
		listStack.Push("test")
	}
}

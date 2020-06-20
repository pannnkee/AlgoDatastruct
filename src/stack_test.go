package src

import "testing"

var stack *ItemStack

func init() {
	stack = newItemStack()
}

func BenchmarkItemStack_Push(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stack.Push("test")
	}
}

func BenchmarkItemStack_Pop(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		stack.Push("test")
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		stack.Pop()
	}
}

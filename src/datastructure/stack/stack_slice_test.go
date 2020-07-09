package stack

import "testing"

var stack *SliceStack

func init() {
	stack = NewStackSlice()
}

func BenchmarkSliceStack_Pop(b *testing.B) {
	for i := 0; i < b.N ; i++ {
		stack.Push("test")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stack.Pop()
	}
}

func BenchmarkSliceStack_Push(b *testing.B) {
	for i := 0; i < b.N ; i++ {
		stack.Push("test")
	}
}


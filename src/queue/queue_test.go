package queue

import "testing"

var queue *Queue

func init() {
	queue = NewQueue()
}

func BenchmarkQueue_Pop(b *testing.B) {


	//b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < b.N; i++ {
				queue.Push("test")
			}

			for i := 0; i < b.N; i++ {
				queue.Pop()
			}
		}
	})
}

func BenchmarkQueue_Push(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < b.N; i++ {
				queue.Push("test")
			}
		}
	})
}

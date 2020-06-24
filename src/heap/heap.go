package heap

import "math"

type MinHeap struct {
	Element []int
}

//创建一个新堆
func NewMinHeap() *MinHeap {
	h := &MinHeap{Element: []int{math.MinInt64}}
	return h
}

func (this *MinHeap) Insert(v int) {
	this.Element = append(this.Element, v)
	i := len(this.Element) - 1
	//上浮
	for ; this.Element[i/2] >v; i/=2 {
		this.Element[i] = this.Element[i/2]
	}
	this.Element[i] = v
}
package asyncgoroutine

import (
	"context"
	"fmt"
	"reflect"
	"sync"
	"time"
)

type GoroutineWorkPool struct {
	wg *sync.WaitGroup
	ch chan interface{}
	chSize int
	goroutineNum int
	task []interface{}
	handleFunc func()
}

func Init(chSize, goroutineNum int, task []interface{}) *GoroutineWorkPool {
	this := new(GoroutineWorkPool)
	this.chSize = chSize
	this.goroutineNum = goroutineNum
	this.task = task

	this.wg = new(sync.WaitGroup)
	this.ch = make(chan interface{}, this.chSize)
	this.wg.Add(len(task))

	return this
}

func (this *GoroutineWorkPool) Produce() {
	//for i := 0; i < 50; i++ {
	//	this.ch <- i
	//}
	for _, v := range this.task {
		this.ch <- v
	}
}

func (this *GoroutineWorkPool) Consume(ctx context.Context) {
	for i := 0; i < this.goroutineNum; i++ {
		go func(i int, ctx context.Context) {
			for  {
				select {
				case task := <-this.ch:
					this.wg.Done()
					fmt.Println("Goroutine:", i, "handle:", task)
				case <-ctx.Done():
					fmt.Println("Cancel:", i)
					return
				default:
					fmt.Println("dont have")
				}
			}
		}(i, ctx)
	}
}

func WorkerInstance() {

	a := []string{"ni","hao","zhong","guo"}
	arg, _ := takeSliceArg(a)

	gp := Init(10, 3, arg)
	// 使用context控制是否停止，适合多级函数传递和控制，并且有超时取消
	ctx, cancel := context.WithCancel(context.Background())
	go gp.Produce()
	go gp.Consume(ctx)

	gp.wg.Wait()
	cancel()
	fmt.Println("Done")
	time.Sleep(time.Second)

}

func PrintHello() {
	fmt.Println("hello")
}

func Print(a []interface{}) {
	for _, v := range a {
		fmt.Println(v)
	}
}

func takeSliceArg(arg interface{}) (out []interface{}, ok bool) {
	slice, success := takeArg(arg, reflect.Slice)
	if !success {
		ok = false
		return
	}
	c := slice.Len()
	out = make([]interface{}, c)
	for i := 0; i < c; i++ {
		out[i] = slice.Index(i).Interface()
	}
	return out, true
}

func takeArg(arg interface{}, kind reflect.Kind) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)
	if val.Kind() == kind {
		ok = true
	}
	return
}

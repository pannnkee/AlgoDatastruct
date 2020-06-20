package main

import (
	"DataStructureGolang/src"
	"sync"
)

func main() {
	//src.ListInstance()
	src.DListNodeInstance()
	mutex := sync.RWMutex{}
	mutex.Lock()
}

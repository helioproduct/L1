package main

import "sync"
import "fmt"

type ConcurrentMap struct {
	data  map[string]string
	taken chan bool
}

func (cm ConcurrentMap) Get(key string) string {
	return cm.data[key]
}

func (cm ConcurrentMap) Add(key, value string) {
	cm.taken <- true
	cm.data[key] = value
	<-cm.taken
}

func NewConcurrentMap() *ConcurrentMap {
	cmPointer := new(ConcurrentMap)
	cmPointer.data = make(map[string]string)
	cmPointer.taken = make(chan bool, 1)
	return cmPointer
}

func main() {

	var wg sync.WaitGroup

	cm := NewConcurrentMap()

	cm.Add("key", "value")

	wg.Add(1)
	go func() {
		cm.Add("key1", "value1")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(cm.Get("key1"))

}

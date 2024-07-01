package main

import "fmt"

type Set map[interface{}]struct{}

func (s Set) Add(key interface{}) {
	s[key] = struct{}{}
}

func (s Set) AddAll(sourceArray []string) {
	for _, str := range sourceArray {
		s.Add(str)
	}
}

func (s Set) Remove(key interface{}) {
	delete(s, key)
}

func (s Set) Has(key interface{}) bool {
	_, ok := s[key]
	return ok
}

func main() {

	sourceArray := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(Set)
	set.AddAll(sourceArray)

	for key, _ := range set {
		fmt.Println(key)
	}
}

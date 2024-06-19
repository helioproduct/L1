package main

import "fmt"

type Set map[interface{}]struct{}

func (s Set) Add(key interface{}) {
	s[key] = struct{}{}
}

func (s Set) Remove(key interface{}) {
	delete(s, key)
}

func (s Set) Has(key interface{}) bool {
	_, ok := s[key]
	return ok
}

func Intersect(leftSet, rightSet Set) Set {
	intersection := make(Set)
	for key, _ := range leftSet {
		if rightSet.Has(key) {
			intersection.Add(key)
		}
	}
	return intersection
}

func main() {
	s1 := make(Set)
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	s2 := make(Set)
	s2.Add(1)
	s2.Add(2)
	s2.Add(4)

	for key, _ := range Intersect(s1, s2) {
		fmt.Println(key)
	}
}

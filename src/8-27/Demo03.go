package main

import (
	"fmt"
	"sort"
)

type Student struct {
	ID   int
	Age  int
	Name string
}

type StudentArr []*Student

func (s StudentArr) Len() int {
	return len(s)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (s StudentArr) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}

// Swap swaps the elements with indexes i and j.
func (s StudentArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	student01 := Student{
		ID:   1,
		Age:  10,
		Name: "a",
	}
	student02 := Student{
		ID:   2,
		Age:  10,
		Name: "b",
	}
	sile := []Student{student01, student02}
	fmt.Println(sile)
	fmt.Println("--------")
	sort.Sort(StudentArr(sile))
	fmt.Println(sile)
}

package main

import (
	"fmt"
)

func main() {
	//testSlice(10)
	slice02 := testSlice02()
	fmt.Println(slice02)
}

func testSlice(lenth int)  {
	sile :=make([] int,lenth)
	for i:=0;i<10 ;i++  {
		sile[i] = i
		fmt.Printf("%p\n",&sile)
	}
}

func testSlice02() ([]int )  {
	s1 := []int{1, 2, 3, 45}
	s2 := make([] int,2)
	copy(s2,s1)
	return s2
}
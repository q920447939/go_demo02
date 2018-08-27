package main

import "fmt"

const (
	ONE = 1 << iota
	TWO
	THREE
)

func main() {
	//	fmt.Printf("I am [%d] years old。",ONE )
	fmt.Printf("%d\n%d\n%d\n", ONE, TWO, THREE)

	c := 1 >> 4

	fmt.Println(c)
}

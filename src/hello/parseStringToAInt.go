package main

import (
	"fmt"
	"time"
)

const (
	timeout  = time.Monday
	timeout2 = 5 * time.Minute
)

func main() {
	i := "10"
	s := fmt.Sprintf(i)
	x := 2
	fmt.Println(s + string(x))

	fmt.Println("---------------")
	fmt.Println("%T", timeout)

}

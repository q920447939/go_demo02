package main

import "fmt"

/**
使用匿名函数
*/

func main() {
	sile := func() []int {
		arr := [3]int{1, 2, 3}
		sile := make([]int, len(arr))
		for e := range arr {
			fmt.Println(e)
			i := arr[e]
			fmt.Println("i:", i)
			sile[e] = e
		}
		return sile
	}()
	fmt.Println(sile)
}

package main

import (
	e "employ"
	"fmt"
)

func main() {
	var sile []e.User
	for i := 0; i < 5; i++ {
		user := e.User{}
		user.ID = i
		user.Name = "王大厨"
		sile = append(sile, user)
	}
	//fmt.Println(sile)
	for idx, v := range sile {
		if v.ID > 1 {
			fmt.Print(sile[idx])
		}
	}

}

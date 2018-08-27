package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	s2 := s1[:len(s1)-1]
	fmt.Printf("%d\n", s2)

	maps := make(map[string]int)
	maps["age"] = 12
	fmt.Println(maps)
	delete(maps, "age")
	maps["age"]++
	fmt.Println(maps)

}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

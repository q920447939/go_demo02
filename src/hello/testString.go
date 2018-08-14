package main

import "fmt"

func main() {
	//testString01()
	s := testRune()
	fmt.Println(s)
}

func testString01()  {
	str := `hello world`
	s1 := str[0:5]
	s2 := str[6:]

	fmt.Println("s1:",s1)
	fmt.Println("s2:",s2)
}

func testRune()(string)  {
	str := `hello world`
	sile01 := []rune(str)
	sile01[0] = '啊'
	str = string(sile01)
	return str
}
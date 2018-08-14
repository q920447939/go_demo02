package main

import "fmt"

func main() {
	//testString01()
	/*s := testRune()
	fmt.Println(s)*/

	map01 := testMap01(nil)
	for k, v := range map01["user"] {
		fmt.Println(k,v)
	}
}

func testString01() {
	str := `hello world`
	s1 := str[0:5]
	s2 := str[6:]

	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
}

func testRune() (string) {
	str := `hello world`
	sile01 := []rune(str)
	sile01[0] = '啊'
	str = string(sile01)
	return str
}

func testMap01(a map[string]map[string]string) (map[string]map[string]string) {
	userMap, ok := a["user"]
	if ok {
		userMap["name"] = `liming  if`
	} else {
		a =make(map[string]map[string]string)
		a["user"] = make(map[string]string)
		a ["user"]["name"] = `liming else`
		a ["user"]["password"] = `pwd else`
	}
	return a
}

package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var node1 *ListNode
	node1.Val = 1

	var node2 *ListNode
	node2.Val = 2

	/*	var node3 *ListNode= new(ListNode)
		node3.Val = 3*/
	node1.Next = node2
	//	node2.Next = node3
	addTwoNumbers(node1, nil)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var sss string
	m1 := getNodeNuM(sss, l1)
	m2 := getNodeNuM(sss, l2)
	m1 = reverseString(m1)
	m2 = reverseString(m2)
	fmt.Println("m1:", m1)
	fmt.Println("m2:", m2)
	return nil
}

func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

func getNodeNuM(str string, l1 *ListNode) string {
	str += strconv.Itoa(l1.Val)
	if l1.Next != nil {
		return getNodeNuM(str, l1.Next)
	}
	return str
}

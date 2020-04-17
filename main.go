package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// n5 := ListNode{Val: 5, Next: nil}
	//n4 := ListNode{Val: 4, Next: nil}
	n3 := ListNode{Val: 3, Next: nil}
	n2 := ListNode{Val: 2, Next: &n3}

	n1 := ListNode{Val: 1, Next: &n2}
	fmt.Println(swapPairs(&n1))
}

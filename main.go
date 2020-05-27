package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// n5 := ListNode{Val: 5, Next: nil}
	//n4 := ListNode{Val: 4, Next: nil}
	// n3 := ListNode{Val: 3, Next: nil}
	// n2 := ListNode{Val: 2, Next: &n3}

	// n1 := ListNode{Val: 1, Next: &n2}
	nums := []int{-1, 2, 1, -4}
	// nums := []int{3, 2, 2, 3}

	// t := [][]string{{"5", "3", ".", ".", "7", ".", ".", ".", "."}, {"6", ".", ".", "1", "9", "5", ".", ".", "."}, {".", "9", "8", ".", ".", ".", ".", "6", "."}, {"8", ".", ".", ".", "6", ".", ".", ".", "3"}, {"4", ".", ".", "8", ".", "3", ".", ".", "1"}, {"7", ".", ".", ".", "2", ".", ".", ".", "6"}, {".", "6", ".", ".", ".", ".", "2", "8", "."}, {".", ".", ".", "4", "1", "9", ".", ".", "5"}, {".", ".", ".", ".", "8", ".", ".", "7", "9"}}
	fmt.Println(threeSumClosest(nums, 1))
}

package MinStack

import (
	"fmt"
	"testing"
)

func Test_Constructor(t *testing.T) {
	action := []string{"MinStack", "push", "push", "push", "getMin", "pop", "top", "getMin"}
	params := [][]int{[]int{}, []int{-3}, []int{-2}, []int{0}, []int{}, []int{}, []int{}, []int{}}

	var obj MinStack
	for index, act := range action {
		if act == "MinStack" {
			obj = Constructor()
		} else if act == "push" {
			obj.Push(params[index][0])
			fmt.Println("---------------------------------------------")
		} else if act == "getMin" {
			fmt.Println("Get Min")

			fmt.Println(obj.GetMin())

			fmt.Println("---------------------------------------------")

		} else if act == "top" {
			fmt.Println("TOP")

			fmt.Println(obj.Top())

			fmt.Println("---------------------------------------------")

		} else {
			fmt.Println("pop")

			obj.Pop()

			fmt.Println("---------------------------------------------")

		}
	}

}

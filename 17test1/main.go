package main

import (
	"fmt"
	"math"
)

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A int, B int) int {
	// write your code in Go 1.4
	start := int(math.Sqrt(float64(A)))
	var result int
	checkPoint1, checkPoint2 := false, false
	for x := start; x <= B; x++ {
		if x*(x+1) < A {
			continue
		}

		if x*(x+1) == A {
			checkPoint1 = true
			result++
			continue
		}

		if x*(x+1) == B {
			checkPoint2 = true
			result++
			break
		}

		result++

	}

	if checkPoint1 && checkPoint2 {
		return result
	}
	return 0

}

func main() {
	fmt.Println(Solution(6, 20))
}

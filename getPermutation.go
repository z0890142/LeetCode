package main

import "fmt"

func getPermutation(n int, k int) string {
	k = k - 1
	x, p := []int{1}, []int{1}
	for i := 2; i <= n; i++ {
		x = append(x, i)
		p = append(p, i*p[len(p)-1])
	}
	fmt.Println(x, p)
	res := []byte{}
	for i := n - 2; i >= 0; i-- {
		pos := k / p[i]
		k = k % p[i]
		fmt.Println(x[pos], byte(x[pos]), byte(x[pos]+'0'))
		res = append(res, byte(x[pos]+'0'))
		x = append(x[:pos], x[pos+1:]...)
	}
	res = append(res, byte(x[0]+'0'))
	return string(res)

}

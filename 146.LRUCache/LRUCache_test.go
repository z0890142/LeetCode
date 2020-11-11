package LRUCache

import (
	"fmt"
	"testing"
)

func Test_Constructor(t *testing.T) {
	action := []string{"LRUCache", "put", "put", "put", "put", "put", "get", "put", "get", "get", "put", "get", "put", "put", "put", "get", "put", "get", "get", "get", "get", "put", "put", "get", "get", "get", "put", "put", "get", "put", "get", "put", "get", "get", "get", "put", "put", "put", "get", "put", "get", "get", "put", "put", "get", "put", "put", "put", "put", "get", "put", "put", "get", "put", "put", "get", "put", "put", "put", "put", "put", "get", "put", "put", "get", "put", "get", "get", "get", "put", "get", "get", "put", "put", "put", "put", "get", "put", "put", "put", "put", "get", "get", "get", "put", "put", "put", "get", "put", "put", "put", "get", "put", "put", "put", "get", "get", "get", "put", "put", "put", "put", "get", "put", "put", "put", "put", "put", "put", "put"}
	params := [][]int{[]int{10}, []int{10, 13}, []int{3, 17}, []int{6, 11}, []int{10, 5}, []int{9, 10}, []int{13}, []int{2, 19}, []int{2}, []int{3}, []int{5, 25}, []int{8}, []int{9, 22}, []int{5, 5}, []int{1, 30}, []int{11}, []int{9, 12}, []int{7}, []int{5}, []int{8}, []int{9}, []int{4, 30}, []int{9, 3}, []int{9}, []int{10}, []int{10}, []int{6, 14}, []int{3, 1}, []int{3}, []int{10, 11}, []int{8}, []int{2, 14}, []int{1}, []int{5}, []int{4}, []int{11, 4}, []int{12, 24}, []int{5, 18}, []int{13}, []int{7, 23}, []int{8}, []int{12}, []int{3, 27}, []int{2, 12}, []int{5}, []int{2, 9}, []int{13, 4}, []int{8, 18}, []int{1, 7}, []int{6}, []int{9, 29}, []int{8, 21}, []int{5}, []int{6, 30}, []int{1, 12}, []int{10}, []int{4, 15}, []int{7, 22}, []int{11, 26}, []int{8, 17}, []int{9, 29}, []int{5}, []int{3, 4}, []int{11, 30}, []int{12}, []int{4, 29}, []int{3}, []int{9}, []int{6}, []int{3, 4}, []int{1}, []int{10}, []int{3, 29}, []int{10, 28}, []int{1, 20}, []int{11, 13}, []int{3}, []int{3, 12}, []int{3, 8}, []int{10, 9}, []int{3, 26}, []int{8}, []int{7}, []int{5}, []int{13, 17}, []int{2, 27}, []int{11, 15}, []int{12}, []int{9, 19}, []int{2, 15}, []int{3, 16}, []int{1}, []int{12, 17}, []int{9, 1}, []int{6, 19}, []int{4}, []int{5}, []int{5}, []int{8, 1}, []int{11, 7}, []int{5, 2}, []int{9, 28}, []int{1}, []int{2, 2}, []int{7, 4}, []int{4, 22}, []int{7, 24}, []int{9, 26}, []int{13, 28}, []int{11, 26}}
	// action := []string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"}
	// params := [][]int{[]int{2}, []int{1, 1}, []int{2, 2}, []int{1}, []int{3, 3}, []int{2}, []int{4, 4}, []int{1}, []int{3}, []int{4}}
	var lRUCache LRUCache
	for index, act := range action {
		if act == "LRUCache" {
			lRUCache = Constructor(params[index][0])
		} else if act == "put" {

			lRUCache.Put(params[index][0], params[index][1])

			fmt.Println("---------------------------------------------")
		} else {
			fmt.Println("Get :", params[index][0])

			fmt.Println(lRUCache.Get(params[index][0]))

			fmt.Println("---------------------------------------------")

		}
	}

}

package main

import (
	"fmt"
	"sort"
)

func solution(arr []int, k int) bool {
	sort.Ints(arr)
	numbersMap := map[int]bool{}
	for i := 0; i < len(arr); i++ {
		if _, ok := numbersMap[k-arr[i]]; ok {
			return true
		}
		numbersMap[arr[i]] = true
	}
	return false
}

func main() {
	arr := []int{10, 15, 1, 3}
	k := 17

	fmt.Println(solution(arr, k))
}

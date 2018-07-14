package main

func solution(arr []int) (res []int) {
	multiplication := 1
	for _, v := range arr {
		multiplication *= v
	}

	for _, v := range arr {
		res = append(res, multiplication/v)
	}

	return res
}

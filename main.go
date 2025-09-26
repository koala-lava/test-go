package main

import "fmt"

func main() {
	a := []int{5, 1, 2, 5}
	b := []int{4, 2, 5, 1, 1, 2}

	aUnique, bUnique := uniq(a, b)
	fmt.Println(aUnique, bUnique)
	fmt.Println(intersect(aUnique, bUnique))
	fmt.Println(union(aUnique, bUnique))
}

func uniq(a []int, b []int) ([]int, []int) {
	return helper([]int{a[0]}, a), helper([]int{b[0]}, b)
}

func helper(result []int, x []int) []int {
	seen := map[int]bool{result[0]: true}
	for _, v := range x[1:] {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

func intersect(a []int, b []int) (result []int) {
	result = []int{}
	for _, v := range a {
		for _, v1 := range b {
			if v == v1 {
				result = append(result, v)
				break
			}
		}
	}
	return
}

func union(a []int, b []int) []int {
	result := append([]int{}, a...)
	for _, v := range b {
		exists := false
		for _, v1 := range result {
			if v == v1 {
				exists = true
				break
			}
		}
		if !exists {
			result = append(result, v)
		}
	}
	return result
}

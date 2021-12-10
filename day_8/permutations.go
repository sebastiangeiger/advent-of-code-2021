package day_8

import "sort"

func less(a []int, b []int) bool {
	if len(a) == len(b) {
		for i := 0; i < len(a); i++ {
			if a[i] > b[i] {
				return false
			} else if a[i] < b[i] {
				return true
			}
		}
		return false
	} else {
		panic("Not the same length")
	}
}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	sort.SliceStable(res, func(i, j int) bool { return less(res[i], res[j]) })
	return res
}

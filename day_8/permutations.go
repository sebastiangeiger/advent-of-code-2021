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

func permutations(input []int) [][]int {
	result := permutationsHelper([]int{}, input)
	sort.SliceStable(result, func(i, j int) bool { return less(result[i], result[j]) })
	return result
}

func permutationsHelper(currentPath []int, remainder []int) [][]int {
	if len(remainder) == 0 {
		return [][]int{currentPath}
	} else {
		paths := [][]int{}
		for i := 0; i < len(remainder); i++ {
			shifted := leftShift(remainder, i)
			newPath := append(currentPath, shifted[0])
			recursiveResults := permutationsHelper(newPath, shifted[1:])
			paths = append(paths, recursiveResults...)
		}
		return paths
	}
}

func leftShift(arr []int, shiftAmount int) []int {
	result := make([]int, len(arr))
	for i := range arr {
		shiftedIndex := (i + shiftAmount) % len(arr)
		result[i] = arr[shiftedIndex]
	}
	return result
}

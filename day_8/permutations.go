package day_8

func permutations(input []int) [][]int {
	return permutationsHelper([]int{}, input)
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

package day_8

import (
	"reflect"
	"testing"
)

func TestApplyMappingForIdentity(t *testing.T) {
	var mapping = map[string]string{
		"a": "a",
		"b": "b",
		"c": "c",
		"d": "d",
		"e": "e",
		"f": "f",
		"g": "g",
	}
	input := "abefg"
	want := input
	result := applyMapping(input, mapping)
	if result != want {
		t.Errorf("applyMapping was wrong, expected '%s' but got '%s'", want, result)
	}
}

func TestApplyMappingForComplicated(t *testing.T) {
	var mapping = map[string]string{
		"a": "b",
		"b": "a",
		"c": "d",
		"d": "c",
		"e": "g",
		"f": "e",
		"g": "f",
	}
	input := "abefg"
	want := "bagef"
	result := applyMapping(input, mapping)
	if result != want {
		t.Errorf("applyMapping was wrong, expected '%s' but got '%s'", want, result)
	}
}

func TestLeftShiftBy0(t *testing.T) {
	input := []int{1, 2, 3, 4}
	want := []int{1, 2, 3, 4}
	result := leftShift(input, 0)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("leftShift was wrong, expected '%v' but got '%v'", want, result)
	}
}

func TestLeftShiftBy2(t *testing.T) {
	input := []int{1, 2, 3, 4}
	want := []int{3, 4, 1, 2}
	result := leftShift(input, 2)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("leftShift was wrong, expected '%v' but got '%v'", want, result)
	}
}

func TestPermutationsWith1(t *testing.T) {
	want := [][]int{[]int{1}}
	result := permutations([]int{1})
	if !reflect.DeepEqual(result, want) {
		t.Errorf("permutations was wrong, expected '%v' but got '%v'", want, result)
	}
}

func TestPermutationsWith2(t *testing.T) {
	want := [][]int{
		[]int{1, 2},
		[]int{2, 1},
	}
	result := permutations([]int{1, 2})
	if !reflect.DeepEqual(result, want) {
		t.Errorf("permutations was wrong, expected '%v' but got '%v'", want, result)
	}
}

func TestPermutationsWith3(t *testing.T) {
	want := [][]int{
		[]int{1, 2, 3},
		[]int{1, 3, 2},
		[]int{2, 1, 3},
		[]int{2, 3, 1},
		[]int{3, 1, 2},
		[]int{3, 2, 1},
	}
	result := permutations([]int{1, 2, 3})
	if !reflect.DeepEqual(result, want) {
		t.Errorf("permutations was wrong, expected '%v' but got '%v'", want, result)
	}
}

func TestPermutationsWith4(t *testing.T) {
	want := [][]int{
		[]int{1, 2, 3, 4},
		[]int{1, 2, 4, 3},
		[]int{1, 3, 2, 4},
		[]int{1, 3, 4, 2},
		[]int{1, 4, 2, 3},
		[]int{1, 4, 3, 2},
		[]int{2, 1, 3, 4},
		[]int{2, 1, 4, 3},
		[]int{2, 3, 1, 4},
		[]int{2, 3, 4, 1},
		[]int{2, 4, 1, 3},
		[]int{2, 4, 3, 1},
		[]int{3, 1, 2, 4},
		[]int{3, 1, 4, 2},
		[]int{3, 2, 1, 4},
		[]int{3, 2, 4, 1},
		[]int{3, 4, 1, 2},
		[]int{3, 4, 2, 1},
		[]int{4, 1, 2, 3},
		[]int{4, 1, 3, 2},
		[]int{4, 2, 1, 3},
		[]int{4, 2, 3, 1},
		[]int{4, 3, 1, 2},
		[]int{4, 3, 2, 1},
	}
	result := permutations([]int{1, 2, 3, 4})
	if !reflect.DeepEqual(result, want) {
		t.Errorf("permutations was wrong, expected '%v' but got '%v'", want, result)
	}
}

func TestPermutationsWith6(t *testing.T) {
	result := permutations([]int{1, 2, 3, 4, 5, 6})
	for i := 1; i < len(result); i++ {
		if reflect.DeepEqual(result[i], result[i-1]) {
			t.Errorf("Found the same value twice %v and %v", result[i], result[i-1])
		}
	}
}

func TestPermutationsWith7(t *testing.T) {
	result := permutations([]int{1, 2, 3, 4, 5, 6, 7})
	problemCount := 0
	for i := 1; i < len(result); i++ {
		if reflect.DeepEqual(result[i], result[i-1]) {
			problemCount += 1
		}
	}
	t.Errorf("Found the same value twice %d times", problemCount)
}

package day_8

import "testing"

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

package helpers

import "testing"

func TestArrayContains(t *testing.T) {
	arr := []string{"str1", "str2", "str3"}

	if !ArrayContains(arr, "str2") {
		t.Errorf("TestArrayContains failed, expected %t, got %t", true, false)
	}
	if ArrayContains(arr, "arr4") {
		t.Errorf("TestArrayContains failed, expected %t, got %t", false, true)
	}
}

func TestArrayEquals(t *testing.T) {
	arr1 := []string{"str1", "str2", "str3"}
	arr2 := []string{"str1", "str2", "str3"}
	arr3 := []string{"str1", "str2", "str3", "str4"}
	arr4 := []string{"str2", "str1", "str3"}

	if !ArrayEquals(arr1, arr2) {
		t.Errorf("TestArrayContains failed, expected %t, got %t", true, false)
	}
	if ArrayEquals(arr1, arr3) {
		t.Errorf("TestArrayContains failed, expected %t, got %t", false, true)
	}
	if ArrayEquals(arr1, arr4) {
		t.Errorf("TestArrayContains failed, expected %t, got %t", false, true)
	}
}

package helpers

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestArrayContains(t *testing.T) {
	arr := []string{"str1", "str2", "str3"}
	require.True(t, ArrayContains(arr, "str2"))
	require.False(t, ArrayContains(arr, "str4"))
}

func TestArrayEquals(t *testing.T) {
	arr1 := []string{"str1", "str2", "str3"}
	arr2 := []string{"str1", "str2", "str3"}
	arr3 := []string{"str1", "str2", "str3", "str4"}
	arr4 := []string{"str2", "str1", "str3"}

	require.True(t, ArrayEquals(arr1, arr2))
	require.False(t, ArrayEquals(arr1, arr3))
	require.False(t, ArrayEquals(arr1, arr4))
}

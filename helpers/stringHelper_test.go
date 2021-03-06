package helpers

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestReplaceForRestParam(t *testing.T) {
	param := map[string]string{
		"description": "first row\n*second row\n[third row]",
	}

	expectDesc := "description=first+row%0A%2Asecond+row%0A%5Bthird+row%5D"
	actual := ReplaceForRestParam(&param)
	require.False(t, expectDesc != actual)

	param["labels"] = strings.Join([]string{"lab-1", "lab 2"}, GetLabelsSeparator())
	actual = ReplaceForRestParam(&param)
	require.True(t, strings.Contains(actual, expectDesc))
	require.True(t, strings.Contains(actual, "labels=lab-1,lab+2"))
}

func TestGetLabelsSeparator(t *testing.T) {
	expect := ","
	actual := GetLabelsSeparator()
	require.Equal(t, expect, actual)
}

func TestFindAllImageLinks(t *testing.T) {
	expect := strings.Join([]string{"image.jpg", "test/image.jpeg", "./test/image.jpg", "image.png"}, ",")
	actual := strings.Join(FindAllImageLinks("test message image.jpg test message, gitlab upload ![abc](/uploads/df1f/image.jpeg) " +
		"and test/image.jpeg, .jpeg, and also ./test/image.jpg and no ![test](/uploads/a0940c63c191c97e471e0b6687da8ee6/test.png)\nimage.png"), ",")
	require.Equal(t, expect, actual)
}
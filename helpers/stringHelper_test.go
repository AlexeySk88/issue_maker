package helpers

import (
	"strings"
	"testing"
)

func TestReplaceForRestParam(t *testing.T) {
	param := map[string]string{
		"description": "first row\n*second row\n[third row]",
	}

	expectDesc := "description=first+row%0A%2Asecond+row%0A%5Bthird+row%5D"
	actual := ReplaceForRestParam(&param)
	if expectDesc != actual {
		t.Errorf("TestReplaceForRestParam failed, expected %s, got %s", expectDesc, actual)
	}

	param["labels"] = strings.Join([]string{"lab-1", "lab 2"}, GetLabelsSeparator())
	actual = ReplaceForRestParam(&param)
	if !strings.Contains(actual, expectDesc) ||
		!strings.Contains(actual, "labels=lab-1,lab+2") {
		t.Errorf("TestReplaceForRestParam failed, got %s", actual)
	}
}

func TestGetLabelsSeparator(t *testing.T) {
	expect := ","
	actual := GetLabelsSeparator()
	if expect != actual {
		t.Errorf("TestGetLabelsSeparator failed, expected %s, got %s", expect, actual)
	}
}

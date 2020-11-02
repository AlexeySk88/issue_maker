package helpers

func ArrayContains(array []string, value interface{}) bool {
	for _, a := range array {
		if a == value {
			return true
		}
	}
	return false
}

func ArrayEquals(arr1 []string, arr2 []string) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for index,_ := range arr1 {
		if arr1[index] != arr2[index] {
			return false
		}
	}

	return true
}

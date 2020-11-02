package helpers

func Contains(array []string, value interface{}) bool {
	for _, a := range array {
		if a == value {
			return true
		}
	}
	return false
}

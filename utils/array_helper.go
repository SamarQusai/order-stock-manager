package utils

func InArrayError(value error, array []error) bool {
	for _, b := range array {
		if b == value {
			return true
		}
	}
	return false
}

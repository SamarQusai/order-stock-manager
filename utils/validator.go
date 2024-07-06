package utils

func IsNull(v interface{}) bool {
	return v == nil || v == ""
}

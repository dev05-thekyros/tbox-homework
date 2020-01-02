package common

// StringInArray Check a string is in string array or not
func StringInArray(value string, values []string) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}

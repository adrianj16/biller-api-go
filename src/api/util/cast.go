package util

func BoolToInt(boolean bool) int {
	if boolean {
		return 1
	} else {
		return 0
	}
}

func IntToBool(integer int) bool {
	if integer == 1 {
		return true
	} else {
		return false
	}
}

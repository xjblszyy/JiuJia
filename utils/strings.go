package utils

func FindString(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func FindInt(intArray []int, val int) bool {
	for _, item := range intArray {
		if item == val {
			return true
		}
	}
	return false
}

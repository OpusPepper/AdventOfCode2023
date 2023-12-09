package utility

func totalIntSlice(intSlice []int) int {
	returnVal := 0

	for _, r := range intSlice {
		returnVal += r
	}

	return returnVal
}

package entities

func CoalesceStrings(strings ...string) string {
	var result string
	for _, oneString := range strings {
		if oneString != "" {
			result = oneString
			break
		}
	}

	return result
}

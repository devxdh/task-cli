package helper

func IsValidTitle(title string) bool {
	if len(title) == 0 {
		return false
	}
	return true
}

func IsValidDescription(description string) bool {
	if len(description) == 0 {
		return false
	}
	return true
}

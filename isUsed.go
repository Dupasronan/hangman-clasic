package pendu

func IsUsed(letter string, usedLetters []string) bool {
	// Checks if a letter is already used
	used := false
	for _, i := range usedLetters {
		if letter == i {
			used = true
		}
	}
	return used
}

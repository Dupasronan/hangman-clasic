package pendu

func IsIn(letter string, word string) bool {
	// Checks if a letter is in a word
	for _, char := range word {
		if string(char) == letter {
			return true
		}
	}
	return false
}

package pendu

import (
	"bufio"
	"fmt"
	"os"
)

func PrintInterface(hiddenWord string, jose int, usedLetters []string) {
	// Opens the file hangman.txt, initialization of the variable hangman
	hangman, _ := os.Open("../hangman.txt")
	defer hangman.Close()
	// Reads hangman.txt
	scanner := bufio.NewScanner(hangman)
	var joseList []string
	// Makes a list with what was read in hangman
	for scanner.Scan() {
		joseList = append(joseList, scanner.Text())
	}
	// Prints the advencement of jose, the hangman
	if jose == 1 {
		for i := (jose - 1) * 7; i < jose*7; i++ {
			fmt.Println(joseList[i])
		}
	} else if jose > 1 {
		for i := (jose-1)*7 + jose - 1; i < (jose-1)*7+jose+7; i++ {
			fmt.Println(joseList[i])
		}
	} else {
		fmt.Println()
	}
	// Prints the used letters
	first := true
	var word string
	for _, i := range usedLetters {
		if first {
			first = false
			word += i
		} else {
			word += ", "
			word += i
		}
	}
	fmt.Println("You used the letters :", word)
	// Prints the hidden word
	fmt.Println(hiddenWord)
}

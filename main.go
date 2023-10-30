package main

import (
	"fmt"
	"os"
	"pendu"
)

func main() {
	fmt.Println("Start the game")

	// Catch the dictionary name
	args := os.Args

	// Verify if there is a dictionary in arguments
	if len(args) == 2 && pendu.IsTxt(args[1]) {

		// Setting of the hidden Word and the word To Find for the game
		hiddenWord, wordToFind := pendu.Start(args[1])

		// Initialisation of Jose our error count
		jose := 0
		usedLetters := []string{}

		// The game loop
		for jose < 10 {

			// Print the interface for the progresion of the hangman and also the letters used
			pendu.PrintInterface(hiddenWord, jose, usedLetters)

			// Catch the user imput
			input := pendu.UserChoice()

			// The case where input is a used letter
			if pendu.IsUsed(input, usedLetters) {
				fmt.Println("This letter has allredy being gest !")

				// The case where the input is in the word To Find
			} else if pendu.IsIn(input, wordToFind) {
				hiddenWord = pendu.Reveal(input, hiddenWord, wordToFind)

				// The case where the input is not in the word To Find
			} else if !pendu.IsIn(input, wordToFind) {
				jose++
				usedLetters = append(usedLetters, input)
				fmt.Println("This letter is not in the word !")
			}

			// The case where the player dosent have more try "lose case"
			if jose == 10 {
				pendu.PrintInterface(hiddenWord, jose, usedLetters)
				fmt.Println("You lose, jose was doomed !\nThe word was :" + wordToFind)
			}

			// The case where the hidden word if find "victory case"
			if hiddenWord == wordToFind {
				fmt.Println("The word find is : " + hiddenWord + ", jose is saved !")

				// Leave the loop because without it the program loops indefinitely
				break
			}
		}
		fmt.Println("You finish the party.\nPlay again ? Y/N")

		// Catch the user answer
		input := pendu.UserChoice()

		// Case where the user starts another party
		if input == "Y" || input == "y" {
			main()

			// Case where the user leave the game
		} else {
			fmt.Println("End of program.")
		}

		// Print where there is no dictionary or the dictionary is not an .txt
	} else {
		fmt.Println("There is no dictionary in arguments. Or the type of the document given in parameters is not an .txt .")
	}
}

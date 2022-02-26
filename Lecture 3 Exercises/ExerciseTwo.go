package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	const oneStepRange int = 1
	const maxGuessedNumber int = 100
	var guessedNumber int = rand.Intn(maxGuessedNumber) + oneStepRange
	const attemptsNumber int = 10
	const minGuessedNumber int = 1

	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)

	fmt.Println(
		"\nWelcome to the game! You need to guess the number the random " +
			"number in range [1,100].\nThe number of your attempts is 10." +
			"\nAfter each answer I tell you 'Your guess was LOW' if you " +
			"are less or 'Your guess was HIGH' if more.")

	for currentAttempt := 0; currentAttempt < attemptsNumber; currentAttempt++ {
		fmt.Printf("\nThe number of remained attempts is '%d'\n", attemptsNumber-currentAttempt)
		fmt.Print("Please, enter your number: ")
		scanner.Scan()
		userNumber, err := strconv.ParseInt(scanner.Text(), 10, 0)

		if err != nil {
			log.Fatal("You need to enter a number correctly\n", err)
		} else if int(userNumber) < minGuessedNumber || int(userNumber) > maxGuessedNumber {
			log.Fatal("You nedd to enter a number in range [1,100]")
		}

		if int(userNumber) < guessedNumber {
			fmt.Println("Oops. Your guess was LOW")
		} else if int(userNumber) > guessedNumber {
			fmt.Println("Oops. Your guess was HIGH")
		} else {
			fmt.Println("You win the game! Congradulations!")
			return
		}
	}
	fmt.Println("\nSorry. You didn't guess my number. It was:", guessedNumber)
}

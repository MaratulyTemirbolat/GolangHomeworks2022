package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func oneInputPart() {
	var totalSum float64 = 0.0
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a number: ")
	input, _ := reader.ReadString('\n')
	inputNumber, err := strconv.ParseFloat(strings.TrimRight(input, "\n"), 64)

	if err != nil {
		fmt.Println("You made a mistake. You need to enter a number in correct format")
		log.Fatal(err)
	}

	totalSum += inputNumber

	fmt.Print("Add more? [y/n] ")
	input, _ = reader.ReadString('\n')

	fmt.Println("Sum is", totalSum)
}

func severalInputsPart() {
	var totalSum float64 = 0.0
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a number: ")
	input, _ := reader.ReadString('\n')
	inputNumber, err := strconv.ParseFloat(strings.TrimRight(input, "\n"), 64)

	if err != nil {
		fmt.Println("You made a mistake. You need to enter a number in correct format")
		log.Fatal(err)
	}

	totalSum += inputNumber
	for {
		fmt.Print("Add more? [y/n] ")
		input, _ = reader.ReadString('\n')

		if strings.TrimRight(input, "\n") != "y" && strings.TrimRight(input, "\n") != "n" {
			fmt.Println("You can enter 'y' or 'n' in lower case only.")
			continue
		} else if strings.TrimRight(input, "\n") == "n" {
			break
		} else {
			fmt.Print("Enter a number: ")
			input, _ = reader.ReadString('\n')
			inputNumber, err = strconv.ParseFloat(strings.TrimRight(input, "\n"), 64)

			if err != nil {
				fmt.Println("You made a mistake. You need to enter a number in correct format")
				continue
			}
			totalSum += inputNumber
		}
	}

	fmt.Println("Sum is", totalSum)
}

func main() {
	// Function with program that's meant to ask the user for a series of numbers and add them together
	//severalInputsPart()

	// Function with program that's meant to ask single number
	//oneInputPart()
}

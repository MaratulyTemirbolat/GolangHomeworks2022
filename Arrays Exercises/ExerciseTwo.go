package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var myArray [100]int
	var valuesRepeatition [3]int

	for curIndex := 0; curIndex < len(myArray); curIndex++ {
		myArray[curIndex] = rand.Intn(3)
		valuesRepeatition[myArray[curIndex]]++
	}

	for curIndex := 0; curIndex < len(valuesRepeatition); curIndex++ {
		fmt.Printf(
			"The repetition of number %d is %d\n",
			curIndex,
			valuesRepeatition[curIndex])
	}

}

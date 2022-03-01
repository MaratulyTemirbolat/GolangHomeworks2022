package main

import (
	"fmt"
)

func main() {
	var weekDays [5]string = [5]string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
	}
	// Option 1
	for index, value := range weekDays {
		fmt.Println(index, value)
	}
	fmt.Println()
	// Option 2
	for curIndex := 0; curIndex < len(weekDays); curIndex++ {
		fmt.Println(curIndex, weekDays[curIndex])
	}
}

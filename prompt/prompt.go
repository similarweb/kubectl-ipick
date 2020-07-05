package prompt

import (
	"fmt"
	"strconv"
)

// InteractiveNumber will prompt a message to the user with number selection
func InteractiveNumber(text string, selectCount int) int {

	// Auto select the first element if we have only one element
	if selectCount == 1 {
		return 1
	}

	var input string
	fmt.Printf("%s: ", text)
	var selectedRowNumber int

	for {
		fmt.Scanln(&input)
		inputNumber, err := strconv.Atoi(input)
		if err == nil && inputNumber != 0 && inputNumber <= selectCount {
			selectedRowNumber = inputNumber
			break
		}
		fmt.Printf("Number must be between %d - %d: ", 1, selectCount)

	}
	return selectedRowNumber
}

// InteractiveText will prompt a message to the user with free text
func InteractiveText(text string) string {

	fmt.Printf("%s: ", text)
	textEnter := ""
	fmt.Scanln(&textEnter)
	return textEnter
}

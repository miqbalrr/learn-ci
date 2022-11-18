package main

import (
	"fmt"
	"learnci/calculator"
)

func main() {
	testID := calculator.Sum(2, 2)
	fmt.Println(testID)
}

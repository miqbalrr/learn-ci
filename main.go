package main

import (
	"fmt"
	"learnci/calculator"
)

func main() {
	testId := calculator.Sum(2, 2)
	fmt.Println(testId)
}

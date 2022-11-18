package calculator

import "errors"

func Sum(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func ThisError() (string, error) {
	return "", errors.New("error")
}

package lab2

import (
	"fmt"
	"strconv"
	"strings"
)

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func PostfixToInfix(input string) (string, error) {
	return converter(strings.Split(input, " "))
}

func converter(input []string) (string, error) {
	var operators = []string{"+", "-", "*", "/", "^"}
	if len(input) < 3 {
		return "", fmt.Errorf("Unable to convert")
	} else if len(input) == 3 {
		var right = input[0]
		var left = input[1]
		var operator = input[2]
		if contains(operators, operator) && isNumeric(left) && isNumeric(right) {
			return fmt.Sprintf("%s %s %s", left, operator, right), nil
		} else {
			return "", fmt.Errorf("Unable to convert")
		}
	}

	var poppedInput = input
	poppedInput = remove(poppedInput, len(poppedInput)-1)
	poppedInput = remove(poppedInput, len(poppedInput)-1)

	var result, error = converter(poppedInput)
	var left = input[len(input)-2]
	var operator = input[len(input)-1]
	if contains(operators, operator) && isNumeric(left) && error == nil {
		return fmt.Sprintf("%s %s (%s)", input[len(input)-2], input[len(input)-1], result), nil
	} else {
		return "", fmt.Errorf("Unable to convert")
	}
}

package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	if strings.TrimSpace(input) == "" {
		return "", fmt.Errorf("invalid input: %w", errorEmptyInput)
	}

	var clearedInput = strings.ReplaceAll(input, " ", "")

	var isFirstOperandNegative = strings.HasPrefix(clearedInput, "-")

	if len(clearedInput) > 1 && (strings.HasPrefix(clearedInput, "+") || strings.HasPrefix(clearedInput, "-")) {
		clearedInput = clearedInput[1:]
	}

	var numberOfOperators = strings.Count(clearedInput, "+") + strings.Count(clearedInput, "-")
	if numberOfOperators != 1 {
		return "", fmt.Errorf("invalid input: %w", errorNotTwoOperands)
	}

	var firstOperand, secondOperand, operator string = "", "", ""

	var operatorIndex = strings.IndexFunc(clearedInput, getOperatorIndex)
	operator = string(clearedInput[operatorIndex])

	firstOperand = clearedInput[0:operatorIndex]

	if operatorIndex < (len(clearedInput) - 1) {
		secondOperand = clearedInput[operatorIndex+1:]
	}

	if len(firstOperand) == 0 || len(secondOperand) == 0 {
		return "", fmt.Errorf("invalid input: %w", errorNotTwoOperands)
	}

	var firstOperandValue, secondOperandValue int = 0, 0

	if firstOperandValue, err = strconv.Atoi(firstOperand); err != nil {
		return "", fmt.Errorf("invalid first operand: %w", err)
	}

	if secondOperandValue, err = strconv.Atoi(secondOperand); err != nil {
		return "", fmt.Errorf("invalid second operand: %w", err)
	}

	if isFirstOperandNegative {
		firstOperandValue = firstOperandValue * -1
	}

	if operator == "+" {
		output = strconv.FormatInt(int64(firstOperandValue+secondOperandValue), 10)
	} else {
		output = strconv.FormatInt(int64(firstOperandValue-secondOperandValue), 10)
	}

	return output, nil
}

func getOperatorIndex(char rune) bool {
	return char == '+' || char == '-'
}

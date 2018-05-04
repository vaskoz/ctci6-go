package twentysix

import (
	"log"
	"strconv"
)

// Calculate take a string representing a computation and returns the result.
func Calculate(input string) float64 {
	operands := []float64{}
	operators := []string{}
	index := 0
	for i, r := range input {
		if r == '*' || r == '+' || r == '-' || r == '/' {
			operands = append(operands, ParseFloat(input[index:i]))
			operators = append(operators, string(r))
			index = i + 1
		}
	}
	operands = append(operands, ParseFloat(input[index:]))
	operandLen, operatorLen := len(operands), len(operators)
	result := 0.0
	for i, o := range operators {
		if o == "*" {
			if result == 0.0 {
				result = operands[i] * operands[i+1]
			} else {
				result *= operands[i+1]
			}
		} else if o == "/" {
			if result == 0.0 {
				result = operands[i] / operands[i+1]
			} else {
				result /= operands[i+1]
			}
		} else {
			if result != 0.0 {
				operands = append(operands, result)
				result = 0.0
			} else {
				operands = append(operands, operands[i])
			}
			operators = append(operators, o)
		}
	}
	operands = append(operands, operands[operandLen-1])
	operands = operands[operandLen:]
	operators = operators[operatorLen:]
	result = operands[0]
	for i, o := range operators {
		if o == "+" {
			result += operands[i+1]
		} else if o == "-" {
			result -= operands[i+1]
		}
	}
	return result
}

// ParseFloat converts a string into a float64.
func ParseFloat(val string) float64 {
	var num float64
	var err error
	if num, err = strconv.ParseFloat(val, 64); err != nil {
		log.Panicf("Got an invalid operand '%v'\n", err)
	}
	return num
}

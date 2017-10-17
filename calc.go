package calc

import (
	"strings"
	"strconv"
)

func ChangeStringtoArray(input string) string{
	elements := strings.Fields(input)

	if len(elements) == 3 {
		return getOperator(elements, input)
	} else {
		return `Wrong elements, please insert only two numbers with operator (+, -, *, /) in between, additionally add space between operator and numbers.`
	}
}

// Retrieve operator from the array
func getOperator(el []string, input string) string{
	operator := strings.Join(el[1:2], "")
	number := strings.Join(el[0:1], "")
	number2 := strings.Join(el[2:3], "")
	result := calculate(toNumber(number), toNumber(number2), operator)
	return result
}

// Make operation and return result
func calculate(num float64, num2 float64, operator string) string {
	switch operator {
		case "+":
			return toString(num + num2)
		case "-":
			return toString(num - num2)
		case "/":
			return toString(num / num2)
		case "*":
			return toString(num * num2)
		default:
			return "Invalid operator: "+operator
	}
}

// Convert string to number
func toNumber(n string) float64 {
	iN, err := strconv.ParseFloat(n, 64)
	if err != nil {
		panic(err)
	}
	return iN
}

// Convert number to string
func toString(iN float64) string {
	return strconv.FormatFloat(iN, 'f', -1, 64)
}

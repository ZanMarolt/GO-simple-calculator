package main

import ( 
	"flag"
	"fmt" 
	"strings"
	"strconv" 
)

func main() {

	input := flag.String("c", "", "Operation you would like to get a result from")

	flag.Parse()

	elements := strings.Fields(*input)

	if len(elements) == 3 { 
		getOperator(elements, *input)
	} else {
		fmt.Println(`Wrong elements, please insert only two numbers with operator (+, -, *, /) in between, additionally add space between operator and numbers.`)
	}

}

// Retrieve operator from the array
func getOperator(el []string, input string){
	operator := strings.Join(el[1:2], "")
	number := strings.Join(el[0:1], "")
	number2 := strings.Join(el[2:3], "")
	result := calculate(toNumber(number), toNumber(number2), operator)
	fmt.Println(input, "=", result)
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

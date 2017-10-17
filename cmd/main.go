package main

import (
  "fmt"
	"flag"
	"go-simple-calculator"
)

func main() {
	input := flag.String("c", "", "Operation you would like to get a result from")

	flag.Parse()

	result := calc.ChangeStringtoArray(*input)

  fmt.Println(*input, "=", result)
}

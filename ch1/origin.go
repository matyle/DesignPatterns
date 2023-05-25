package ch1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// implement calculator
// 1. input from std: 1 + 2 * 3
// 2. std output: 7

func OriginCalculator() {
	// read numbers from std
	// read operator from std
	// calculate
	// print result to std
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input numbers: ")
	numstr1, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Read number1 error: ", err)
		return
	}
	numstr2, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Read number2 error: ", err)
		return
	}
	fmt.Println("Please input operator: ")
	operator, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Read operator error: ", err)
		return
	}
	fmt.Println("num1: ", numstr1, "num2: ", numstr2, "operator: ", operator)

	// calculate
	num1, err := strconv.ParseFloat(numstr1, 64)
	if err != nil {
		fmt.Println("Parse number1 error: ", err)
		return
	}
	num2, _ := strconv.ParseFloat(numstr2, 64)
	if err != nil {
		fmt.Println("Parse number2 error: ", err)
		return
	}
	var result float64
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	}
	fmt.Println("result: ", result)
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Calc calculates reverse Polish notation expression
func Calc(inputStream io.Reader, outputStream io.Writer) (resultErr error) {
	defer func() {
		if err := recover(); err != nil {
			resultErr = fmt.Errorf("recover: %v", err)
		}
	}()

	reader := bufio.NewReader(inputStream)

	var stack []int
	expression, _ := reader.ReadString('\n')

	for _, char := range expression {
		switch char {
		case ' ':
		case '=':
			if len(stack) != 1 {
				panic("Incorrect input format")
			}
			fmt.Fprintf(outputStream, "Result = %d\n", stack[len(stack)-1])
			return nil
		case '+':
			if len(stack) >= 2 {
				op1, op2 := stack[len(stack)-2], stack[len(stack)-1]
				stack = stack[:len(stack)-2]
				res := op1 + op2
				stack = append(stack, res)
			} else {
				panic("Incorrect input format")
			}
			break
		case '-':
			if len(stack) >= 2 {
				op1, op2 := stack[len(stack)-2], stack[len(stack)-1]
				stack = stack[:len(stack)-2]
				res := op1 - op2
				stack = append(stack, res)
			} else {
				panic("Incorrect input format")
			}
			break
		case '*':
			if len(stack) >= 2 {
				op1, op2 := stack[len(stack)-2], stack[len(stack)-1]
				stack = stack[:len(stack)-2]
				res := op1 * op2
				stack = append(stack, res)
			} else {
				panic("Incorrect input format")
			}
			break
		case '/':
			if len(stack) >= 2 {
				op1, op2 := stack[len(stack)-2], stack[len(stack)-1]
				stack = stack[:len(stack)-2]
				res := op1 / op2
				stack = append(stack, res)
			} else {
				panic("Incorrect input format")
			}
			break
		default:
			if value, err := strconv.Atoi(string(char)); err == nil {
				stack = append(stack, value)
			} else {
				panic("Incorrect input format")
			}
		}
	}
	return nil
}

func main() {
	err := Calc(os.Stdin, os.Stdout)
	if err != nil {
		fmt.Printf("unexpected error %v", err)
	}
}

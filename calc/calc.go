package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func calc(inputStream io.Reader, outputStream io.Writer) (resultErr error) {
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
			fmt.Fprintf(outputStream, "Result = %d\n", stack[len(stack)-1])
			return nil
		case '+':
			op1, op2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			res := op1 + op2
			stack = append(stack, res)
			break
		case '-':
			op1, op2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			res := op1 - op2
			stack = append(stack, res)
			break
		case '*':
			op1, op2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			res := op1 * op2
			stack = append(stack, res)
			break
		case '/':
			op1, op2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			res := op1 / op2
			stack = append(stack, res)
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
	err := calc(os.Stdin, os.Stdout)
	if err != nil {
		fmt.Printf("unexpected error %v", err)
	}
}

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

	stack := [10000]int{}
	sp := 0
	expression, _ := reader.ReadString('\n')

	for _, char := range expression {
		switch char {
		case ' ':
		case '=':
			fmt.Fprintf(outputStream, "Result = %d\n", stack[sp-1])
			sp--
			return nil
		case '+':
			stack[sp-2] = stack[sp-2] + stack[sp-1]
			sp--
			break
		case '-':
			stack[sp-2] = stack[sp-2] - stack[sp-1]
			sp--
			break
		case '*':
			stack[sp-2] = stack[sp-2] * stack[sp-1]
			sp--
			break
		case '/':
			stack[sp-2] = stack[sp-2] / stack[sp-1]
			sp--
			break
		default:
			if value, err := strconv.Atoi(string(char)); err == nil {
				stack[sp] = value
				sp++
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

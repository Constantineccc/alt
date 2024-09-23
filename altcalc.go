package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите значение:")
	expression, _ := reader.ReadString('\n')
	expression = strings.TrimSpace(expression)

	result, err := evaluateExpression(expression)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func evaluateExpression(expression string) (string, error) {

	expression = strings.ReplaceAll(expression, " ", "")

	var operator string
	var index int
	for i, char := range expression {
		if char == '+' || char == '-' || char == '*' || char == '/' {
			operator = string(char)
			index = i
			break
		}
	}

	if operator == "" {
		return "", fmt.Errorf("значение не соответсвует условию задачи")
	}

	leftPart := expression[:index]
	rightPart := expression[index+1:]

	var leftString string
	var leftNumber int
	var err error
	if leftPart[0] == '"' && leftPart[len(leftPart)-1] == '"' {
		leftString = leftPart[1 : len(leftPart)-1]
	} else {
		leftNumber, err = strconv.Atoi(leftPart)
		if err != nil {
			return "", fmt.Errorf("левая часть должна быть строкой или целым числом")
		}
	}

	var rightString string
	var rightNumber int
	if operator == "+" || operator == "-" {
		if rightPart[0] == '"' && rightPart[len(rightPart)-1] == '"' {
			rightString = rightPart[1 : len(rightPart)-1]
		} else {
			return "", fmt.Errorf("правая часть должна быть строкой для операций + и -")
		}
	} else {
		rightNumber, err = strconv.Atoi(rightPart)
		if err != nil || rightNumber < 1 || rightNumber > 10 {
			return "", fmt.Errorf("правая часть должна быть целым числом от 1 до 10 для операций * и /")
		}
	}

	var result string
	switch operator {
	case "+":
		if leftString != "" {
			result = leftString + rightString
		} else {
			result = strconv.Itoa(leftNumber) + rightString
		}
	case "-":
		if leftString != "" {
			result = strings.Replace(leftString, rightString, "", 1)
		} else {
			return "", fmt.Errorf("операция вычитания не поддерживается для целых чисел")
		}
	case "*":
		if leftString != "" {
			result = strings.Repeat(leftString, rightNumber)
		} else {
			result = strconv.Itoa(leftNumber * rightNumber)
		}
	case "/":
		if leftString != "" {
			if len(leftString) < rightNumber {
				return "", fmt.Errorf("результатом деления будет пустая строка")
			}
			result = leftString[:len(leftString)/rightNumber]
		} else {
			result = strconv.Itoa(leftNumber / rightNumber)
		}
	default:
		return "", fmt.Errorf("неподдерживаемая операция")
	}

	if len(result) > 40 {
		result = result[:40] + "..."
	}

	return result, nil
}

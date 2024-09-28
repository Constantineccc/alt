package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var input string
	fmt.Print("Введите выражение: ")
	fmt.Scanln(&input)

	result, err := calculate(input)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func calculate(expression string) (string, error) {

	pattern := `^"([^"]{1,10})"\s*([+\-*/])\s*("([^"]{1,10})"|(\d+))$`
	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(expression)

	if len(match) == 0 {
		return "", fmt.Errorf("Неверный формат выражения")
	}

	str1 := match[1]
	operation := match[2]
	operand := match[3]

	// Проверяем длину строк
	if len(str1) > 10 {
		return "", fmt.Errorf("Длина первой строки превышает 10 символов")
	}


	var str2 string
	var num int
	if strings.HasPrefix(operand, `"`) {
		str2 = operand[1 : len(operand)-1]
		if len(str2) > 10 {
			return "", fmt.Errorf("Длина второй строки превышает 10 символов")
		}
	} else {
		var err error
		num, err = strconv.Atoi(operand)
		if err != nil || num < 1 || num > 10 {
			return "", fmt.Errorf("Число должно быть в диапазоне от 1 до 10")
		}
	}


	var result string
	switch operation {
	case "+":
		result = str1 + str2
	case "-":
		result = strings.Replace(str1, str2, "", 1)
	case "*":
		result = strings.Repeat(str1, num)
	case "/":
		result = str1[:len(str1)/num]
	default:
		return "", fmt.Errorf("Неподдерживаемая операция")
	}


	if len(result) > 40 {
		result = result[:40] + "..."
	}

	return `"` + result + `"`, nil
}

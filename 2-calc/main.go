package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func UserScan() {
	var operation string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		operation = scanner.Text()
		if operation == "AVG" || operation == "SUM" || operation == "MED" {
			break
		} else {
			fmt.Println("Ошибка, значение должно быть AVG, SUM, MED")
		}
	}
	fmt.Println("Введите желаемые числа через запятую c пробелом")
	for {
		ok := scanner.Scan()
		if !ok {
			fmt.Println("Ошибка ввода")
		} else {
			break
		}
	}
	text := scanner.Text()
	numbers := strings.Split(text, ",")
	sum := 0
	for _, number := range numbers {
		number_int, _ := strconv.Atoi(number)
		sum += number_int
	}
	switch operation {
	case "AVG":
		fmt.Println(sum / len(numbers))
	case "SUM":
		fmt.Println(sum)
	case "MED":
		fmt.Println(numbers[len(numbers)/2])
	}
}

func main() {
	UserScan()
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var mapRomNum = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

var mapArabNum = map[int]string{
	100: "C",
	90:  "XC",
	50:  "L",
	40:  "XL",
	10:  "X",
	9:   "IX",
	5:   "V",
	4:   "IV",
	1:   "I",
}

var mySort = [9]int{
	100,
	90,
	50,
	40,
	10,
	9,
	5,
	4,
	1,
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Введите значение: ")
	_ = scanner.Scan()
	example := scanner.Text()

	fmt.Println(calculate(example))
}

func calculate(example string) string {

	text := strings.Split(strings.TrimSpace(strings.ToUpper(example)), "")
	operator := ""

	for _, val := range text {
		if val == "+" || val == "-" || val == "*" || val == "/" {
			operator += string(val)
			break
		}
	}

	strNum := strings.Split(strings.TrimSpace(strings.ToUpper(example)), operator)

	if len(strNum) < 2 {
		panic("Выдача паники, так как строка не является математической операцией.")
	}
	if len(strNum) > 2 {
		panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}

	firstNum, err1 := strconv.Atoi(strings.TrimSpace(strNum[0]))
	sacondNum, err2 := strconv.Atoi(strings.TrimSpace(strNum[1]))
	if err1 == nil && err2 == nil {
		if firstNum <= 0 || sacondNum <= 0 {
			panic(fmt.Sprint("Выдача паники, некорректное число."))
		}
		if firstNum > 10 || sacondNum > 10 {
			panic(fmt.Sprint("Выдача паники, некорректное число."))

		}

		return strconv.Itoa(count(firstNum, sacondNum, operator))

	}
	if err1 != nil && err2 != nil {
		firstNum = romToInt(strings.TrimSpace(strNum[0]))
		sacondNum = romToInt(strings.TrimSpace(strNum[1]))
		if firstNum > 10 || sacondNum > 10 {
			panic(fmt.Sprint("Выдача паники, некорректное число."))
		}
		return arabToRom(count(firstNum, sacondNum, operator))
	}
	if err1 != nil || err2 != nil {
		panic(fmt.Sprint("Выдача паники, так как используются одновременно разные системы счисления."))
	}

	panic(fmt.Sprint("Некоректное выражение"))
}

func romToInt(arg string) int {

	_, isPresent := mapRomNum[arg]
	if isPresent {
		return mapRomNum[arg]
	}
	return 100

}
func count(firstNum, sacondNum int, operator string) int {

	result := 0
	switch operator {
	case "+":
		result = firstNum + sacondNum
	case "-":
		result = firstNum - sacondNum
	case "*":
		result = firstNum * sacondNum
	case "/":
		result = firstNum / sacondNum

	}

	return result

}
func arabToRom(num int) string {
	if num < 1 {
		panic("Выдача паники, так как в римской системе нет отрицательных чисел.")
	}
	var res string = ""
	var n int = num
	for n > 0 {
		for _, val := range mySort {
			if val <= n {
				res += mapArabNum[val]
				n -= val
				break
			}

		}
	}
	return res

}

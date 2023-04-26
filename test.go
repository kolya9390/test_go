package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanNumerals = map[string]int{
	"I": 1,
	"II": 2,
	"III": 3,
	"IV": 4,
	"V": 5,
	"VI": 6,
	"VII": 7,
	"VIII": 8,
	"IX": 9,
	"X": 10,
}

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) int {
	return a / b
}

func romanToArabic(romanNumeral string) (int, error) {
	romanNumerals := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
	}

	var result int
	for i := 0; i < len(romanNumeral); i++ {
		value := romanNumerals[rune(romanNumeral[i])]
		if i+1 < len(romanNumeral) {
			nextValue := romanNumerals[rune(romanNumeral[i+1])]
			if value < nextValue {
				value = -value
			}
		}
		result += value
	}

	return result, nil
}

func arabicToRoman(arabicNumeral int) (string, error) {
	
	romanNumerals := []struct {
		Value  int
		Symbol string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var result string
	for _, numeral := range romanNumerals {
		for arabicNumeral >= numeral.Value {
			result += numeral.Symbol
			arabicNumeral -= numeral.Value
		}
	}

	return result, nil
}

// Проверка принадлежности к римским цифрам.
func isRomanNumeral(numeral string) bool {
	_, ok := romanNumerals[numeral]
	return ok
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение, пример 1 + 1 или I + I: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	

	op1_op_op2 := strings.Split(text, " ")

	if len(op1_op_op2) != 3 {

		fmt.Print("Error: Вы вели неверный формат выражения.")
		return
	}

	operand1Str := op1_op_op2[0]
	operator := op1_op_op2[1]
	operand2Str := op1_op_op2[2]


	if operand1Str == ""|| operator == "" || operand2Str == ""{
		fmt.Print("Error")
		return
	}


	operand1, err := strconv.Atoi(operand1Str)
	if err != nil {
		// Если операнд1 не является числом, пытаемся преобразовать в римское число
		operand1, err = romanToArabic(operand1Str)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	operand2, err := strconv.Atoi(operand2Str)
	if err != nil {
		// Если операнд2 не является числом, пытаемся преобразовать в римское число
		operand2, err = romanToArabic(operand2Str)
		if err != nil {
			fmt.Println(err)
			return
		}
	}


	// Проверка вводимых цифр на принодлежность промежутку от 1 до 10
	if (operand1 > 10 || operand1 < 1) || (operand2 > 10 || operand2 < 1) {

		fmt.Println("Ошибка: вы ввели числа не попадющие в промежуток из условия, " + 
		"указывайте числа в промежутке [1:10].")
		return
	}

	if (isRomanNumeral(operand1Str) && ! isRomanNumeral(operand2Str)) || (! isRomanNumeral(operand1Str) && isRomanNumeral(operand2Str)) {

		fmt.Print("Error: используются одновременно разные системы счисления.")
		return
	}

	// Выполнение операции в зависимости от оператора
	var result int
	switch operator {
	case "+":
		result = add(operand1, operand2)
	case "-":
		result = subtract(operand1, operand2)
	case "*":
		result = multiply(operand1, operand2)
	case "/":
		result = divide(operand1, operand2)
	default:
		fmt.Println("Ошибка: Неверный оператор")
		return
	}


	if ! isRomanNumeral(operand1Str) && ! isRomanNumeral(operand2Str) {
	
	// Вывод результата
	fmt.Println("Результат: (арабские цифры)", result)
	return
	}
	if  isRomanNumeral(operand1Str) &&  isRomanNumeral(operand2Str) {
		if result < 1 {

			fmt.Print("Ошибка: Результат для римских чисел не может быть отрицательным или равен нулю.")
			return
		}
	// Преобразование результата в римское число, если операнды были римскими числами
	romanResult, err := arabicToRoman(result)
	if err == nil {
		fmt.Println("Результат (римские цифры):", romanResult)
	}
	}
}
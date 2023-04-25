package main

import (
	"fmt"
	"strconv"
)

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
		'D': 500,
		'M': 1000,
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
	if arabicNumeral <= 0 {
		return "", fmt.Errorf("Ошибка: Римское число не может быть меньше или равно нулю")
	}

	romanNumerals := []struct {
		Value  int
		Symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
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

func main() {
	var operand1Str string
	var operand2Str string
	var operator string
	fmt.Print("Введите выражение: ")
	fmt.Scanln(&operand1Str, &operator, &operand2Str)

	if operand1Str == "" || operator == "" || operand2Str == "" {
		fmt.Println("Ошибка: Неверный формат выражения 2 ")
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

	// Проверка вводимых арабских цифр на принодлежность промежутку от 1 до 10
	/// if (operand1 > 10 || operand1 < 1) || (operand2 > 10 || operand2 < 1) {

	//	fmt.Println("Erorr")
	//	return
	//}

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

	// Вывод результата
	fmt.Println("Результат:", result)

	// Преобразование результата в римское число, если операнды были римскими числами
	romanResult, err := arabicToRoman(result)
	if err == nil {
		fmt.Println("Результат (римские цифры):", romanResult)
	}
}
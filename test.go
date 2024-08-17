package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanToArabic = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

var arabicToRoman = []string{
	"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX",
	"X", "XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX",
	"XX", "XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX",
	"XXX", "XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX",
	"XL", "XLI", "XLII", "XLIII", "XLIV", "XLV", "XLVI", "XLVII", "XLVIII", "XLIX",
	"L", "LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII", "LIX",
	"LX", "LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX",
	"LXX", "LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX",
	"LXXX", "LXXXI", "LXXXII", "LXXXIII", "LXXXIV", "LXXXV", "LXXXVI", "LXXXVII", "LXXXVIII", "LXXXIX",
	"XC", "XCI", "XCII", "XCIII", "XCIV", "XCV", "XCVI", "XCVII", "XCVIII", "XCIX",
	"C",
}

func romanToInt(roman string) int {
	if value, exists := romanToArabic[roman]; exists {
		return value
	}
	panic("Invalid Roman numeral")
}

func intToRoman(num int) string {
	if num <= 0 || num > 100 {
		panic("Roman numeral out of range")
	}
	return arabicToRoman[num]
}

func calculate(a int, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("Invalid operator")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение (например, 3 + 5 или VI * III):")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Определяем, используются ли арабские или римские числа
	isRoman := false
	for _, char := range input {
		if _, exists := romanToArabic[string(char)]; exists {
			isRoman = true
			break
		}
	}

	// Разделяем строку на части
	parts := strings.Fields(input)
	if len(parts) != 3 {
		panic("Invalid input format")
	}

	aStr, operator, bStr := parts[0], parts[1], parts[2]

	var a, b int
	if isRoman {
		a = romanToInt(aStr)
		b = romanToInt(bStr)
	} else {
		var err error
		a, err = strconv.Atoi(aStr)
		if err != nil || a < 1 || a > 10 {
			panic("Invalid Arabic numeral")
		}
		b, err = strconv.Atoi(bStr)
		if err != nil || b < 1 || b > 10 {
			panic("Invalid Arabic numeral")
		}
	}

	result := calculate(a, b, operator)

	if isRoman {
		if result <= 0 {
			panic("Roman numeral result is less than 1")
		}
		fmt.Println(intToRoman(result))
	} else {
		fmt.Println(result)
	}
}

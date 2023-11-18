package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Калькулятор структура
type Calculator struct{}

// Метод для выполнения операции сложения
func (c Calculator) add(a, b int) int {
	return a + b
}

// Метод для выполнения операции вычитания
func (c Calculator) subtract(a, b int) int {
	return a - b
}

// Метод для выполнения операции умножения
func (c Calculator) multiply(a, b int) int {
	return a * b
}

// Метод для выполнения операции деления
func (c Calculator) divide(a, b int) int {
	return a / b
}

// Метод для преобразования римских чисел в арабские
func romanToArabic(roman string) (int, error) {
	romanNumerals := map[rune]int{'I': 1, 'V': 5, 'X': 10}

	var result int
	prevValue := 0

	for _, r := range roman {
		value, found := romanNumerals[r]
		if !found {
			return 0, fmt.Errorf("неверный символ в римском числе")
		}

		result += value
		if prevValue < value {
			result -= 2 * prevValue
		}

		prevValue = value
	}

	return result, nil
}

// Метод для выполнения операции в зависимости от введенной строки
func (c Calculator) performOperation(input string) (int, error) {
	// Разделение строки на операнды и оператор
	re := regexp.MustCompile(`\s*([-+*/IVXLCDMivxlcdm]+)\s*([-+*/])\s*([-+*/IVXLCDMivxlcdm]+)\s*`)
	matches := re.FindStringSubmatch(input)

	// Проверка соответствия формату
	if len(matches) != 4 {
		return 0, fmt.Errorf("неверный формат математической операции")
	}

	// Преобразование операндов в числа
	operand1, err := strconv.Atoi(matches[1])
	if err != nil {
		// Если не удалось преобразовать в арабское число, пробуем римское
		operand1, err = romanToArabic(matches[1])
		if err != nil {
			return 0, fmt.Errorf("ошибка преобразования первого операнда")
		}
	}

	operand2, err := strconv.Atoi(matches[3])
	if err != nil {
		// Если не удалось преобразовать в арабское число, пробуем римское
		operand2, err = romanToArabic(matches[3])
		if err != nil {
			return 0, fmt.Errorf("ошибка преобразования второго операнда")
		}
	}

	// Выбор операции в зависимости от оператора
	switch matches[2] {
	case "+":
		return c.add(operand1, operand2), nil
	case "-":
		return c.subtract(operand1, operand2), nil
	case "*":
		return c.multiply(operand1, operand2), nil
	case "/":
		// Проверка деления на ноль
		if operand2 == 0 {
			return 0, fmt.Errorf("деление на ноль недопустимо")
		}
		return c.divide(operand1, operand2), nil
	default:
		return 0, fmt.Errorf("недопустимый оператор")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	calculator := Calculator{}

	for {
		// Чтение строки из консоли
		fmt.Print("Input: ")
		scanner.Scan()
		input := scanner.Text()

		// Проверка строки на соответствие условиям задачи
		if strings.TrimSpace(input) == "" {
			fmt.Println("Вывод ошибки, так как строка не является математической операцией.")
			break
		}

		// Выполнение операции
		result, err := calculator.performOperation(input)
		if err != nil {
			fmt.Println("Вывод ошибки:", err)
			break
		}

		// Вывод результата
		fmt.Println("Output:", result)
	}
}

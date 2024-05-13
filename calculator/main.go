package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func f_sum(a, b float64) float64 {
	return a + b
}

func f_dif(a, b float64) float64 {
	return a - b
}

func f_mult(a, b float64) float64 {
	return a * b
}

func f_div(a, b float64) float64 {
	return a / b
}

func main() {

	var a float64
	var b float64
	var oper string
	var result float64
	scanner := bufio.NewScanner(os.Stdin)

	// Проверка первого введенного числа
	for {
		fmt.Print("Введите первое число: ")
		scanner.Scan()
		input := scanner.Text()
		val, err := strconv.ParseFloat(input, 64)

		if err != nil {
			fmt.Println("Некорректное число. Повторите попытку.")
		} else {
			a = val
			break
		}
	}
	fmt.Printf("Вы ввели число: %f\n", a)

	// Проверка введенной операции
	for {
		fmt.Print("Выберите операцию: (+, -, *, /): ")
		fmt.Scanf("%s", &oper)

		if oper == "+" || oper == "-" || oper == "*" || oper == "/" {
			break
		} else {
			fmt.Println("Некорректная операция. Пожалуйста, используйте символы +, -, * или /")
		}
	}

	// Очистка входного буфера
	scanner.Scan()

	// Проверка второго числа
	for {
		fmt.Print("Введите второе число: ")
		scanner.Scan()
		input := scanner.Text()
		val, err := strconv.ParseFloat(input, 64)

		if err != nil {
			fmt.Println("Некорректное число. Пожалуйста, введите числовое значение.")
		} else if oper == "/" && val == 0 {
			fmt.Println("Ошибка: деление на ноль невозможно.")
		} else {
			b = val
			break
		}
	}
	fmt.Printf("Вы ввели второе число: %f\n", b)

	if oper == "+" {
		result = f_sum(a, b)
	} else if oper == "-" {
		result = f_dif(a, b)
	} else if oper == "*" {
		result = f_mult(a, b)
	} else if oper == "/" {
		result = f_div(a, b)
	}

	fmt.Println("Полученный результат: ", result)

}

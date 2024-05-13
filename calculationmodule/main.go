package main

import (
	"calculationmodule/fact"
	"flag"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {

	// Получение значения числа n из командной строки
	args := os.Args
	n := int64(len(args) - 1)
	fmt.Println("Количество переданных аргументов командной строки:", len(args)-1)

	// Получение флага '-log (bool)'
	var Logging bool
	flag.BoolVar(&Logging, "log", false, "Enable logging")
	flag.Parse()

	// Устанавливаем форматтер лога для бибилиотеки
	logrus.SetFormatter(&logrus.TextFormatter{})
	// Устанавливаем уровень логирования INFO
	logrus.SetLevel(logrus.InfoLevel)
	result := fact.Calculate(n, Logging)

	logrus.Info("Factorial of ", n, " is: ", result)

}

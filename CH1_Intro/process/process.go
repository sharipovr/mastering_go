package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Недостаточно агрументов")
		return
	}

	var total, nInts, nFloats int
	invalid := make([]string, 0)

	for _, k := range arguments[1:] {

		// Это целое число?
		_, err := strconv.Atoi(k)
		if err == nil {
			total++
			nInts++
			continue
		}

		// Это число с плавающей точкой?
		_, err = strconv.ParseFloat(k, 64)
		if err == nil {
			total++
			nFloats++
			continue
		}

		// Тогда это некорректное значение
		invalid = append(invalid, k)
	}

	fmt.Println("#read:", total, "#ints:", nInts, "#floats:", nFloats)
	if len(invalid) > total {
		fmt.Println("Слишком много некорректного ввода:", len(invalid))
		for _, s := range invalid {
			fmt.Println(s)
		}
	}
}

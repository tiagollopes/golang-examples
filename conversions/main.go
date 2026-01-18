package main

import (
	"fmt"
	"strconv"
)

func main() {
	// ======================
	// String -> Int
	// ======================
	strInt := "42"
	intValue, err := strconv.Atoi(strInt)
	if err != nil {
		fmt.Println("Erro ao converter string para int:", err)
	} else {
		fmt.Println("String para int:", intValue)
	}

	// ======================
	// Int -> String
	// ======================
	number := 100
	strNumber := strconv.Itoa(number)
	fmt.Println("Int para string:", strNumber)

	// ======================
	// String -> Float
	// ======================
	strFloat := "3.14"
	floatValue, err := strconv.ParseFloat(strFloat, 64)
	if err != nil {
		fmt.Println("Erro ao converter string para float:", err)
	} else {
		fmt.Println("String para float:", floatValue)
	}

	// ======================
	// Float -> String
	// ======================
	floatNumber := 2.71828
	strFromFloat := strconv.FormatFloat(floatNumber, 'f', 2, 64)
	fmt.Println("Float para string:", strFromFloat)

	// ======================
	// String -> Bool
	// ======================
	strBool := "true"
	boolValue, err := strconv.ParseBool(strBool)
	if err != nil {
		fmt.Println("Erro ao converter string para bool:", err)
	} else {
		fmt.Println("String para bool:", boolValue)
	}
}

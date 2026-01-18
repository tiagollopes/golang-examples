package main

import "fmt"

func main() {
	// Inteiros com sinal
	var a int8 = -128
	var b int16 = -32768
	var c int32 = -2147483648
	var d int64 = -9223372036854775808

	// Inteiros sem sinal
	var e uint8 = 255
	var f uint16 = 65535
	var g uint32 = 4294967295
	var h uint64 = 18446744073709551615

	// Tipos dependentes da arquitetura
	var i int = 42
	var j uint = 42

	// Ponto flutuante
	var k float32 = 3.14
	var l float64 = 3.14159265359

	// String e boolean
	var m string = "Go is awesome"
	var n bool = true

	fmt.Println("int8:", a)
	fmt.Println("int16:", b)
	fmt.Println("int32:", c)
	fmt.Println("int64:", d)

	fmt.Println("uint8:", e)
	fmt.Println("uint16:", f)
	fmt.Println("uint32:", g)
	fmt.Println("uint64:", h)

	fmt.Println("int:", i)
	fmt.Println("uint:", j)

	fmt.Println("float32:", k)
	fmt.Println("float64:", l)

	fmt.Println("string:", m)
	fmt.Println("bool:", n)
}

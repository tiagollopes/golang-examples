package main

import "fmt"

// Criando sub-slices (sublistas)
func main() {

	numbers := []int{10, 20, 30, 40, 50}

	// Sub-slice: início incluso, fim exclusivo
	sub1 := numbers[1:4] // 20, 30, 40
	sub2 := numbers[:3]  // 10, 20, 30
	sub3 := numbers[2:]  // 30, 40, 50

	fmt.Println("Slice original:", numbers)
	fmt.Println("Sub-slice 1:", sub1)
	fmt.Println("Sub-slice 2:", sub2)
	fmt.Println("Sub-slice 3:", sub3)
}


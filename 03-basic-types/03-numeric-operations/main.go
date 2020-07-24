package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b = 4, 5
	var res1 = (a + b) * (a + b) / 2

	a++

	b += 10
	//a, b = 1, 2
	fmt.Printf("%d, %d\n", a, b)
	var res2 = a ^ b //xor 按位异或，相同则0不同则1

	var r = 3.5

	var res3 = math.Pi * r * r

	fmt.Printf("res1 : %v, res2 : %v, res3 : %v\n", res1, res2, res3)
}

package main

import "fmt"

func main() {
	//默认int ,默认float64
	var myInt8 int8 = 97

	var myInt = 1200

	var myUint uint = 500

	var myHexNumber = 0xFF  // Use prefix '0x' or '0X' for declaring hexadecimal numbers
	var myOctalNumber = 034 // Use prefix '0' for declaring octal numbers

	var myFloat32 float32 = 4.5
	var myFloat = 9.12 // // Type inferred as `float64` (the default type for floating-point numbers)

	fmt.Printf("%d, %T, %d, %T, %d, %T, %#x, %T, %#o, %T, %f, %T, %f, %T\n", myInt8, myInt8, myInt, myInt, myUint, myUint, myHexNumber, myHexNumber, myOctalNumber, myOctalNumber, myFloat32, myFloat32, myFloat, myFloat)
}

package main

import "fmt"

func main() {
	//定义变量
	var myStr string = "Hello"
	var myInt int = 100
	var myFloat float64 = 54.453
	fmt.Println(myStr, myInt, myFloat)

	//定义多个变量
	var (
		employeeID          int    = 5
		firstName, lastName string = "Satoshi", "Nakamoto"
	)
	fmt.Println(employeeID, firstName, lastName)

	//段变量定义
	name := "Rajeev Singh"
	age, salary, isProgrammer := 35, 50000.0, true
	fmt.Println(name, age, salary, isProgrammer)
}

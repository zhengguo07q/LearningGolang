package main

import "fmt"

func main() {
	//定义一个未定义类型的变量并反射得到类型
	name := "Rajeev Singh"
	fmt.Printf("Variable 'name' is of type %T\n", name)

	//定义多个未定义类型变量
	var firstName, lastName, age, salary = "John", "Maxwell", 28, 50000.0

	fmt.Printf("firstName: %T, lastName: %T, age: %T, salary: %T\n",
		firstName, lastName, age, salary)
}

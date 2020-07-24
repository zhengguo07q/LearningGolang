package main

import "fmt"

func main() {
	//与变量定义一样，只是var改成const 可以不标记类型，也可以一次定义多个，可以使用const()
	const myFavLanguage = "Python"
	const sunRisesInTheEast = true

	const country, code = "India", 91

	const (
		employeeID string  = "E101"
		salary     float64 = 50000.0
	)

	const typedInt int = 1234
	const typedStr string = "Hi"

	fmt.Println(myFavLanguage, sunRisesInTheEast, country, code, employeeID, salary, typedInt, typedStr)
}

package main

import (
	"fmt"
)

func main() {
	a, b := 12.3, 342.2
	c := avg(a, b)
	fmt.Printf("(%f+%f)/2=%f", a, b, c)
}

//普通的函数定义
func avg(x float64, y float64) float64 {
	return (x + y) / 2
}

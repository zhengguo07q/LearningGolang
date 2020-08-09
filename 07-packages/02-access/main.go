package main

import (
	"02-access/model"
	"fmt"
)

func main() {
	p := model.NewPerson("smith")
	p.SetAge(18)
	p.SetSal(5000)
	fmt.Println(p)
	fmt.Println(p.Name, " age =", p.GetAge(), " sal = ", p.GetSal())
}

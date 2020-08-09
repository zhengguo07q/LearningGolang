/*
Package model 人模型

变量或者函数
小写字母开头，则是代表私有
大写开头则是其他包可访问
*/
package model

import "fmt"

type person struct {
	Name string
	age  int //其它包不能直接访问
	sal  float64
}

// NewPerson 工厂模式创建函数
func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("年龄范围不正确")
	}
}

func (p *person) GetAge() int {
	return p.age
}

func (p *person) SetSal(sal float64) {
	if sal >= 3000 && sal <= 30000 {
		p.sal = sal
	} else {
		fmt.Println("薪水范围不正确")
	}
}

func (p *person) GetSal() float64 {
	return p.sal
}

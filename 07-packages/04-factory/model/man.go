package model

import (
	"04-factory/base"
	"fmt"
)

// Man
type Man struct {
}

// 实现Class接口
func (c *Man) Do() {
	fmt.Println("man do")
}

func init() {
	// 在启动时注册类1工厂
	base.Register("man", func() base.Class {
		return new(Man)
	})
}

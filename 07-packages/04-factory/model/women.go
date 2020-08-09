package model

import (
	"04-factory/base"
	"fmt"
)

// Women
type women struct {
}

// 实现Class接口
func (c *women) Do() {
	fmt.Println("women do")
}

func init() {
	// 在启动时注册类1工厂
	base.Register("women", func() base.Class {
		return new(women)
	})
}

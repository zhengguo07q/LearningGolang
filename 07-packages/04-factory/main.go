/*
package main

利用init机制做自动类注册
*/
package main

import (
	"04-factory/base"
	_ "04-factory/model" // 匿名引用cls1包, 自动注册
)

func main() {
	c1 := base.Create("man")
	c1.Do()
	c2 := base.Create("women")
	c2.Do()
	c3 := base.Create("class")
	c3.Do()
}

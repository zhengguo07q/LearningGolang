/*
package Main

init的调用顺序
*/
package main

import (
	"03-init/service"
	"fmt"
)

func main() {
	fmt.Println("main.go main")
	service.HandleReq()
}

func init() {
	fmt.Println("main.go init")
}

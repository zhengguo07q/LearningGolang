package service

import "fmt"

// HandleReq 处理请求
func HandleReq() {
	fmt.Println("service - http.go Handle Request")
}

// init 处理请求
func init() {
	fmt.Println("service - http.go init")
}

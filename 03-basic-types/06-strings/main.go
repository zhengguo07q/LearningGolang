package main

import "fmt"

func main() {
	// 普通字符串，不能包含新的行，但是可以使用各种转义字符
	var website = "\thttps://www.callicoder.com\t\n"

	// 非处理的字符串，严格里面的内容，有什么就输出什么，可以包含自己的格式，不能转义
	var siteDescription = `\t\tCalliCoder is a programming blog where you can find
                           practical guides and tutorials on programming languages, 
                           web development, and desktop app development.\t\n`

	fmt.Println(website, siteDescription)
}

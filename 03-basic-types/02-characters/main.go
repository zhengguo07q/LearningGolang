package main

import "fmt"

func main() {
	//单字节字符 byte unicode用rune 输出%c 和%U
	var myByte byte = 'a'
	var myRune rune = '♥'

	fmt.Printf("%c = %d  type=%T and %c = %U type=%T\n", myByte, myByte, myByte, myRune, myRune, myRune)
}

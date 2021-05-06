package main

import (
	"fmt"
)

func main() {
	var t string = "我是测试字符串"
	var r string
	r = Reverse(t)
	fmt.Println("我是测试字符串")
	fmt.Println(r)
}

func Reverse(s string) string{
	var str []rune = []rune(s)
	fmt.Println(str)
	for i, j := 0, len(str) -1; i < j; i ,j = i +1, j - 1 {
		str[i], str[j] = str[j], str[i]
	}
	fmt.Println(str)
	return string(str)
}

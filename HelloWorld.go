package main

import "fmt"

var (
	aa = 3
	ss = "kkk"
	bb = true
)

func variableZeroValue() {
	var a int
	var b string
	fmt.Printf("%d %q\n", a, b)
}
func variableInitialValue() {
	var a int = 3
	var b string = "黄英杰"
	fmt.Printf("%d %s\n", a, b)
}

func variableTypeDeduction() {
	var a, b, c, d = 3, 4, true, "def"
	fmt.Println(a, b, c, d)
}

func variableShorter() {
	a, b, c, d := 3, 4, true, "def"
	fmt.Println(a, b, c, d)
}

func main() {
	fmt.Println("Hello World!")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(ss)
}

package sample01

import "fmt"

func HelloWorld(s string) string {
	return "hello world, " + s
}

var ExportedVarWithoutComment = 1 // TODO: revert

func Greet() {
	fmt.Println("Hello, World!")
}

var vetError = fmt.Sprintf("", 1, 2, 3) // TODO: revert

package sample01

import "fmt"

func HelloWorld(s string) string {
	var y int
	y = "aaa" // TODO
	return "hello world, " + s
}

var ExportedVarWithoutComment = 1 // TODO: revert

func Greet() {
	fmt.Println("Hello, World!")
}

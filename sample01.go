package sample01

import "fmt"

type Connpass struct {
	ResultsReturned  int `json:results_returned`
	ResultsAvailable int `json:results_available`
	ResultsStart     int `json:results_start`
}

func HelloWorld(s string) string {
	return "hello world, " + s
}

var ExportedVarWithoutComment = 1 // TODO: revert

func Greet() {
	fmt.Println("Hello, World!")
}

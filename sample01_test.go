package sample01

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	actual := HelloWorld("hoge")
	expected := "aaa, hoge"
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

package main

import (
	"fmt"
	"test_project/setup"
	_ "test_project/setup"
	_ "test_project/sp"
)

func init() {
	fmt.Println("init main.go test_project")
}
func main() {
	setup.Hello()
	fmt.Println("main main.go test_project")
	fmt.Println("play with init function")
}

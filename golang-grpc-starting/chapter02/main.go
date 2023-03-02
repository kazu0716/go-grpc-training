package main

import "fmt"

func main() {
	res := hello("hoge")
	fmt.Printf(res)
}

func hello(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

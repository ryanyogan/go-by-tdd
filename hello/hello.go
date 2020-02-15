package main

import "fmt"

// Hello returns the string "Hello, world"
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("Ryan"))
}

package main

import "fmt"

const englishHelloPrefix = "Hello, "

// Hello returns the classic Hello World string
func Hello(name string) string {

	if name == "" {
		return englishHelloPrefix + "World" + "!"
	}
	return englishHelloPrefix + name + "!"
}

func main() {
	fmt.Println(Hello("world"))
}

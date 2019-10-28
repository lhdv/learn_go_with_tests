package main

import "fmt"

const englishHelloPrefix = "Hello, "

const spanishLang = "Spanish"
const spanishHelloPrefix = "Hola, "

// Hello returns the classic Hello World string
func Hello(name, lang string) string {

	if name == "" {
		name = "World"
	}

	if lang == spanishLang {
		return spanishHelloPrefix + name + "!"
	}

	return englishHelloPrefix + name + "!"
}

func main() {
	fmt.Println(Hello("world", ""))
}

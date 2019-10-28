package main

import "fmt"

const englishHelloPrefix = "Hello, "

const spanishLang = "Spanish"
const spanishHelloPrefix = "Hola, "

const frenchLang = "French"
const frenchHelloPrefix = "Bonjour, "

const portugueseLang = "Portuguese"
const portugueseHelloPrefix = "Olá, "

// Hello returns the classic Hello World string
func Hello(name, lang string) string {

	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch lang {
	case spanishLang:
		prefix = spanishHelloPrefix
	case frenchLang:
		prefix = frenchHelloPrefix
	case portugueseLang:
		prefix = portugueseHelloPrefix
	}

	return prefix + name + "!"
}

func main() {
	fmt.Println(Hello("world", ""))
}

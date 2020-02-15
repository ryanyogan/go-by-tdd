package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const frenchHelloPrefix = "Bonjour, "
const spanishHelloPrefix = "Hola, "
const englishHelloPrefix = "Hello, "

// Hello takes a name and language as strings, returns "Hello, [name]" in
// in the language provided, if none given, return will default to English.
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greeetingPrefix(language) + name
}

func greeetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("Ryan", ""))
}

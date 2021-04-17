package main

import "fmt"

const arabic = "Arabic"
const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const arabicHelloPrefix = "مرحبا، "
const spanishHelloPrefix = "Hola, "

func main() {
	fmt.Println(Hello("World", ""))
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case arabic:
		prefix = arabicHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

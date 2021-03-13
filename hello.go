package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const kituba = "Kituba"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const kitubaHelloPrefix = "Mbote, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case kituba:
		prefix = kitubaHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("Tater", ""))
}

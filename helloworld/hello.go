package main

import "fmt"

const (
	spanish            = "Spanish"
	french             = "French"
	catalan            = "Catalan"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	catalanHelloPrefix = "Bon dia, "
)

// Greets people in different languages
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

// Returns the greeting prefix
func greetingPrefix(language string) (prefix string) {
	// 'prefix' is the variable to return
	// it's already created as a string
	// zero value as ""
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case catalan:
		prefix = catalanHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	// Return an empty value 'prefix'
	return
}

func main() {
	fmt.Println(Hello("", ""))
}

package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello"
const spanishHelloPrefix = "Hola"
const frenchHelloPrefix = "Bonjour"

func composeGreeting(helloPrefix string, name string) string {
	return helloPrefix + ", " + name
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == spanish {
		return composeGreeting(spanishHelloPrefix, name)
	}
	if language == french {
		return composeGreeting(frenchHelloPrefix, name)
	}
	return composeGreeting(englishHelloPrefix, name)
}

func main() {
	fmt.Println(Hello("Matheus", ""))
}

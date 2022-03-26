package main

import "fmt"

const spanish = "Spanish"
const englishHelloPrefix = "Hello"
const spanishHelloPrefix = "Hola"

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

	return composeGreeting(englishHelloPrefix, name)
}

func main() {
	fmt.Println(Hello("Matheus", ""))
}

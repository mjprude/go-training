package main

import "fmt"

func localizedHello(language string) string {
	switch language {
	case "French":
		return "Bonjour"
	case "Spanish":
		return "Hola"
	case "English":
		return "Hello"
	}
	return ""
}

// Hello returns a friendly greeting
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return localizedHello(language) + ", " + name
}

func main() {
	fmt.Println(Hello("Michael", "English"))
}

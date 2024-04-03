package main

import (
	"fmt"
)

const ( //grouped together all the constants
	spanish               = "Spanish"
	french                = "French"
	marathi               = "Marathi"
	englishGreetingPrefix = "Hello, "
	spanishGreetingPrefix = "Hola, "
	frenchGreetingPrefix  = "Bonjour, "
	marathiGreetingPrefix = "Namaskar, "
)

func Hello(name string, language string) string { //public function as name starts with capital letter
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) { // function name starting with small letter are private
	switch language {
	case marathi:
		prefix = marathiGreetingPrefix
	case french:
		prefix = frenchGreetingPrefix
	case spanish:
		prefix = spanishGreetingPrefix
	default:
		prefix = englishGreetingPrefix
	}
	return // (prefix string) is a named return so whatever is store in it will be returned just like return prefix
}
func main() {
	fmt.Println(Hello("Nikhil", ""))
}

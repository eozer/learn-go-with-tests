// First exercise from:
// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world
package main

import "fmt"

const spanish = "es"
const french = "fr"
const enHelloPrefix = "Hello "
const esHelloPrefix = "Hola "
const frHelloPrefix = "Bonjour "

// Hello returns "Hello World" used by hello-world_test.go
func Hello(name string, language string) string {
	// NOTE: There are no func overloading, optional or default arguments in golang yet?
	if name == "" {
		name = "World"
	}
	return getGreetings(language) + name + "!"
}

// NOTE: we returned a "named" variable implicitly. This wasnt declared explicitly by a dev in
//  a scope, but go understands it. Moreover, we did not need to "return prefix".
func getGreetings(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = esHelloPrefix
	case french:
		prefix = frHelloPrefix
	default:
		prefix = enHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", ""))
}

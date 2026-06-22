package main

import "fmt"

// Greet returns a friendly greeting for the given name. An empty name
// defaults to "world".
func Greet(name string) string {
	if name == "" {
		name = "world"
	}
	return fmt.Sprintf("Hello, %s!", name)
}

func main() {
	fmt.Println(Greet("Bito"))
}

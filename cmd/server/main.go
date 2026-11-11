package main

import "fmt"

func main() {
	fmt.Println(greet("Gopher"))
}

func greet(name string) string {
	return fmt.Sprintf("Hello %s!\n", name)
}

package main

import (
	"fmt"
	"hello_people/greetings"
)

var (
	name     string
	language string
	greeting string
	err      error
)

func main() {
	fmt.Print("What is your name? ")
	fmt.Scanf("%s", &name)
	fmt.Print("What language do you speak? ")
	fmt.Scanf("%s", &language)

	greeting, err = greetings.Greet(name, language)
	if err == nil {
		fmt.Println(greeting)
	} else {
		fmt.Printf("I'm sorry %s, I do not speak %s.\n", name, language)
	}

}

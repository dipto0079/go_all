package main

import (
	"fmt"
)

type User struct {
	Name  string
	Email string
	age   int
}

func main() {
	rob := User{"rob", "rob@lco.dev", 40}
	fmt.Printf("%v\n", rob)
	fmt.Printf("%v\n", rob.age)

	var sam = new(User)
	sam.Name = "sam"
	sam.Email = "sam@lco.dev"
	sam.age = 26
	fmt.Printf("%v\n", sam)

	var tobby = &User{"tobby", "toby@loc.dev", 23}
	fmt.Printf("%v\n", tobby)

}

//Functions
package main

import (
	"fmt"
)

func main() {
	superman()
	result := multiplyme(2, 5)
	fmt.Println(result)
	myresult, mylength, myname := adder(3, 6, 8, 4, 2, 5)
	fmt.Println(myresult, mylength, myname)
}
func superman() {
	fmt.Println("i am superman")
}
func multiplyme(v1 int, v2 int) int {
	return v1 * v2
}

func adder(valuse ...int) (int, int, string) {
	sum := 0
	for i := range valuse {
		sum = sum + valuse[i]
	}
	langht := len(valuse)
	name := "just for fun"
	return sum, langht, name
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// var myString string
	// fmt.Scanln(&myString)
	// fmt.Println(myString)

	// var name string = "Dipto"
	// var a, b = fmt.Println(name)
	// fmt.Println(a, b)

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter you full name:")
	// mayname, _ := reader.ReadString('\n')
	// fmt.Println(mayname)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Entser Your reating: ")
	myrating, _ := reader.ReadString('\n')
	mynumRating, _ := strconv.ParseFloat(strings.TrimSpace(myrating), 64)
	fmt.Println(mynumRating + 80)
}

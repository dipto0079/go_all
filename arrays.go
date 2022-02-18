package main

import "fmt"

func main() {
	// var lifebooster float64 = 99.2
	// lifeboosterRef := &lifebooster

	// fmt.Println(lifebooster)
	// fmt.Println(*lifeboosterRef)
	var numbers[3] string
	numbers[0]="uno"
	numbers[1]="dos"
	numbers[2]="tres"

	fmt.Println(numbers,"\n")

	var colors = [4]string{
		"roje","gris","azul","black"}

	fmt.Println(colors)
	fmt.Println(colors[2])
	fmt.Println(len(colors))
}

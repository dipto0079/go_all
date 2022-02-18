package main

import (
	"fmt"
	"sort"
)

func main() {
	var things = []string{"maleta", "ropa", "reloj"}
	//fmt.Println(things)

	things = append(things, "bolso")
	fmt.Println(things)

	things = append(things[2:])
	fmt.Println(things)

	heros := make([]string, 3, 3)
	heros[0] = "thor"
	heros[1] = "ironman"
	heros[2] = "spiderman"

	fmt.Println(heros)

	heros = append(heros, "deadpool")
	fmt.Println(heros)
	fmt.Println(cap(heros))

	myints := []int{4, 7, 1, 3, 5, 4}

	isSroted := sort.IntsAreSorted(myints)

	fmt.Println(isSroted)

}

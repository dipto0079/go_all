package main

import "fmt"

func main() {
	score := make(map[string]int)
	score["hitesh"] = 89
	fmt.Println(score)

	score["josh"] = 21
	score["roy"] = 50
	score["joy"] = 40
	score["dipto"] = 80
	score["test"] = 90
	fmt.Println(score)

	getRonScore := score["roy"]
	fmt.Println(getRonScore)

	delete(score, "test")
	fmt.Println(score)

	for K, V := range score {
		fmt.Println("score of %v is %v\n", K, V)
	}

}

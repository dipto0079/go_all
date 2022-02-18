package main

import "fmt"

func main() {
	start := 1
	things := []string{"jony", "joy", "test", "kaj hosa", "all ok"}
	fmt.Println(things)

	for i := 0; i < 10; i++ {
		fmt.Println(i + start)
	}

	for i := 0; i < len(things); i++ {
		fmt.Println(things[i])
	}

	for i := range things {
		fmt.Println(things[i])
	}

	for start < 100 {
		start += start
		if start == 32 {
			//continue
			//break
			goto lcolabel
		}
		fmt.Println("start is now:", start)

	}
lcolabel:
	fmt.Println("Learn")
}

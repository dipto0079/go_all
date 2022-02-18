package main

import (
	"fmt"
	"strconv"
)

func main() {

	a := 5
	//b := 5.6
	b := "5"
	//d, err := strconv.Atoi(b)
	d := strconv.Itoa(a)
	//c := float64(a) + b
	//c := a + int(b)
	//if err != nil {
	//	log.Fatal(err)
	//}
	c := d + b

	fmt.Println("Result:", c)

	//s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//res := sum(s...)
	//nxtInt := seqInt()
	//fmt.Println("Result:", nxtInt())
	//fmt.Println("Result:", nxtInt())
	//
	//newInt := seqInt()
	//fmt.Println("Result:", newInt())
	//
	res := newFunc(sum)
	fmt.Println("Result:", res)

	res = newFunc(mult)
	fmt.Println("Multiplication Result: ", res)
}

func newFunc(s func(nums []int) int) int {
	c := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	return s(c)
}

func seqInt() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)

}

func sum(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	//for i := 0; i < len(nums); i++ {
	//	total += nums[i]
	//}
	return total
}

func mult(nums []int) int {
	total := 1
	for _, num := range nums {
		total *= num
	}
	return total
}

package main

import (
	"fmt"
	"math"
)

//type Person struct {
//	Name    string
//	Age     int
//	weight  float64
//	IsAdult bool
//}

//func New(name string, age int, weight float64) *Person {
//	isAdult := age > 18
//	return &Person{
//		Name:    name,
//		Age:     age,
//		weight:  weight,
//		IsAdult: isAdult,
//	}
//}

//func (p *Person) SetName(name string) {
//	p.Name = name
//}
//
//func (p *Person) GetName() string {
//	return p.Name
//}

type geometry interface {
	area() float64
	perim() float64
}
type Rect struct {
	Height float64
	Width  float64
}

func (r Rect) area() float64 {
	return r.Height * r.Width
}

func (r Rect) perim() float64 {
	return 2 * (r.Height + r.Width)
}

type Circle struct {
	Radis float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.Radis, 2)
}

func (c Circle) perim() float64 {
	return 2 * math.Pi * c.Radis
}

func getArea(g geometry) float64 {
	return g.area()
}

func getPerim(g geometry) float64 {
	return g.perim()
}

func main() {
	var r Rect = Rect{Height: 5, Width: 10}
	fmt.Println("Area:", getArea(r))
	fmt.Println("Perim:", getPerim(r))

	var c Circle = Circle{Radis: 5}
	fmt.Println("Circle Area:", getArea(c))
	fmt.Println("Circle Perim:", getPerim(c))

	//p := Person{}
	//p.SetName("dipto")
	//fmt.Println(p.GetName())
}

package main

import (
	"fmt"
	"os"
)

type shape interface {
	getArea() float64
}
type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

func main() {
	//s := square{sideLength: 10}
	//t := triangle{base: 10, height: 10}
	//printArea(s)
	//printArea(t)
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	bs := make([]byte, 99999)
	f.Read(bs)
	fmt.Println(string(bs))
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return (t.base * t.height) / 2
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

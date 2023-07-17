package main

import (
	"fmt"
)

func main() {
	ynr := User{"Yogendra", "asdf", true, 23}
	fmt.Println(ynr)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

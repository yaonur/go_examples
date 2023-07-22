package main

import "fmt"

func Maps() {
	var mp map[string]int = map[string]int{"apple": 5, "pear": 6, "banana": 7}
	// mp2 := make(map[string]int)
	fmt.Println(mp)
	mp["apple"] = 9000
	fmt.Println(mp["apple"])
	mp["orange"] = 100
	fmt.Println(mp)
	delete(mp, "apple")
	fmt.Println(mp)
	val, ok := mp["orange"]
	fmt.Println("val:", val, "ok:", ok)
}

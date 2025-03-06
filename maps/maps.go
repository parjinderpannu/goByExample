package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("m: ", m)

	delete(m, "k2")
	fmt.Println("m: ", m)

	clear(m)
	fmt.Println("m: ", m)

	_, isKeyPresent := m["k2"]
	fmt.Println(isKeyPresent)

	n := map[string]int{"foo": 7, "bar": 9}
	fmt.Println(n)

	n1 := map[string]int{"foo": 7, "bar": 9}
	fmt.Println(n1)
	if maps.Equal(n, n1) {
		fmt.Println("n == n1")
	}

}

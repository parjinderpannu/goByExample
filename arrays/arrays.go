package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("var a [5]int value: ", a)

	a[4] = 100
	fmt.Println("a[4]=100 a=", a)
	fmt.Println("a[4] = ", a[4])

	b := [5]int{11, 12, 13, 14, 15}
	fmt.Println("b = ", b)

	b = [...]int{41, 21, 44, 67, 99}
	fmt.Println("b:", b)

	var twoD [2][3]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println(twoD)

	twod2 := [2][3]int{
		{7, 8, 9},
		{32, 76, 39},
	}
	fmt.Println(twod2)

}

package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("Write", i, " as ")
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	fmt.Println("Today is ", time.Now().Weekday())
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's Weekend")
	default:
		fmt.Println("Weekday")
	}

	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Before noon")
	case t.Hour() > 12:
		fmt.Println("Afternoon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		default:
			fmt.Printf("don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(4)
	whatAmI("hello")

}

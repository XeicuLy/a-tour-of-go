package main

import (
	"fmt"
	"strconv"
)

func typeChangeFn() {
	var x int = 1
	xx := float64(x)
	fmt.Printf("%T %v %f\n", xx, xx, xx)

	var y float64 = 3.14
	yy := int(y)
	fmt.Printf("%T %v %d\n", yy, yy, yy)

	var s string = "123"
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("%T %v %d\n", i, i, i)
	}

	h := "Hello World"
	fmt.Println(string(h[1]))
}

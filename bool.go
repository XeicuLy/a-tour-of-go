package main

import "fmt"

func boolFn() {
	t, f := true, false
	fmt.Printf("%T %v %t\n", t, 1, t)
	fmt.Printf("%T %v %t\n", f, 0, f)

	fmt.Println(true && true)   // true
	fmt.Println(true && false)  // false
	fmt.Println(false && true)  // false
	fmt.Println(false && false) // false

	fmt.Println(true || true)   // true
	fmt.Println(true || false)  // true
	fmt.Println(false || true)  // true
	fmt.Println(false || false) // false

	fmt.Println(!true)  // false
	fmt.Println(!false) // true
}

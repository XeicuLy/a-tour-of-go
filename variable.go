package main

import "fmt"

// var i int = 1
// var f64 float64 = 1.2
// var s string = "test"
// var t, f bool = true, false

// var (
// 	i   int
// 	f64 float64
// 	s   string
// 	t   bool
// 	f   bool
// )

var (
	i   int     = 1
	f64 float64 = 1.2
	s   string  = "test"
	t   bool    = true
	f   bool    = false
)

func foo() {
	xi := 1
	xf64 := 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)
}

func variable() {
	fmt.Println(i, f64, s, t, f)
	foo()
}

package main

import "fmt"

func variable() {
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
	fmt.Println(i, f64, s, t, f)
}

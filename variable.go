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
	xi = 2 // 変数xiに値を代入する
	var xf64 float32 = 1.2 // 変数宣言の際に型を指定することもできる
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)
	fmt.Printf("%T\n", xf64) // %Tは変数の型を出力するフォーマット指定子(printfの引数に渡す変数の型を出力する)
	fmt.Printf("%T\n", xi)   // 変数xiの型を出力する
}

func variable() {
	fmt.Println(i, f64, s, t, f)
	foo()
}

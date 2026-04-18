package main

import "fmt"

// 一番最初に実行される。(main関数よりも先に実行される)
func init() {
	fmt.Println("Init function called")
}

func buzz() {
	fmt.Println("Buzz")
}

func main() {
	buzz()
	fmt.Println("Hello, World!", "TEST TEST")
}

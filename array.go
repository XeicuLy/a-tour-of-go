package main

import "fmt"

func arrayFn() {
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	var b [2]int = [2]int{100, 200}
	// b = append(b, 300) // 配列はサイズが固定されているため、appendは使用できない
	fmt.Println(b)

	var b2 []int = []int{100, 200} // スライスを使用することで、可変長の配列を作成できる
	b2 = append(b2, 300)           // スライスはサイズが可変のため、appendを使用できる
	fmt.Println(b2)
}

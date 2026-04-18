package main

import "fmt"

func sliceFn() {
	n := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(n)
	fmt.Println(n[0])   // 1
	fmt.Println(n[2:4]) // [3 4]
	fmt.Println(n[:3])  // [1 2 3]
	fmt.Println(n[3:])  // [4 5 6]
	fmt.Println(n[:])   // [1 2 3 4 5 6]

	n[2] = 100
	fmt.Println(n) // [1 2 100 4 5 6]

	var board = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(board) // [[1 2 3] [4 5 6] [7 8 9]]

	n = append(n, 7, 8, 9)
	fmt.Println(n) // [1 2 100 4 5 6 7 8 9]
}

func makeFn() {
	n := make([]int, 3, 5)                                    // 長さ3、容量5のスライスを作成
	fmt.Println(n)                                            // [0 0 0]'
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n) // len=3 cap=5 value=[0 0 0]
	n = append(n, 1, 2)                                       // スライスに要素を追加
	fmt.Println(n)                                            // [0 0 0 1 2]
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n) // len=5 cap=5 value=[0 0 0 1 2]
	n = append(n, 3)                                          // スライスに要素を追加（容量を超えるため、新しい配列が作成される）
	fmt.Println(n)                                            // [0 0 0 1 2 3]
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n) // len=6 cap=10 value=[0 0 0 1 2 3]

	c := make([]int, 3) // 長さ3、容量3のスライスを作成
	for i := range 3 {
		c = append(c, i)
		fmt.Println(c)
	}
	fmt.Println(c)

	d := make([]int, 0, 3) // 長さ0、容量3のスライスを作成
	for i := range 3 {
		d = append(d, i)
		fmt.Println(d)
	}
	fmt.Println(d)
}

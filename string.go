package main

import "fmt"

func stringFn() {
	str := "Hello, World!"
	fmt.Println(str)
	fmt.Println("Hello" + " World!")
	fmt.Println("Hello World"[0])         // 文字列はバイトのスライスとして扱われるため、インデックスでアクセスすると対応するバイト値が返される
	fmt.Printf("%c\n", "Hello World"[0])  // %cは文字を出力するフォーマット指定子
	fmt.Println(string("Hello World"[0])) // CASTして文字列に変換することもできる

	fmt.Println(`Test
	
	Test`) // バックスラッシュをエスケープせずにそのまま文字列に含めることができる
}

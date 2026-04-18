package main

import "fmt"

func mapFn() {
	m := map[string]int{
		"apple":  100,
		"banana": 200,
	}
	fmt.Println(m["apple"])  // 100
	fmt.Println(m["banana"]) // 200

	m["orange"] = 300
	fmt.Println(m["orange"]) // 300

	delete(m, "banana")
	fmt.Println(m["banana"]) // 0（存在しないキーにアクセスすると、ゼロ値が返される）

	if value, ok := m["banana"]; ok {
		fmt.Println(value) // 存在する場合の処理
	} else {
		fmt.Println("banana does not exist") // 存在しない場合の処理
	}

	m2 := make(map[string]int) // 空のマップを作成
	m2["grape"] = 400
	fmt.Println(m2)
	fmt.Println(m2["grape"]) // 400

}

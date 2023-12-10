package main

import "fmt"

//マップ
//keyと対応するvalueを設定することで便利なデータ構造を実現している

func main() {
	var m = map[string]int{"A": 100, "B": 200}
	fmt.Println(m)

	m2 := map[string]int{"A": 100, "B": 200}
	fmt.Println(m2)

	//以下のように宣言することもできる
	m3 := map[int]string{
		1: "A",
		2: "B",
	}
	fmt.Println(m3)

	//からのマップ
	m4 := make(map[int]string)
	fmt.Println(m4)

	//値の追加
	m4[1] = "JAPAN"
	m4[2] = "USA"

	fmt.Println(m4)

	//値の取り出し

	fmt.Println(m["A"])
	fmt.Println(m4[2])
	fmt.Println(m4[3])
	//mapの三番はない、キーがない場合は初期値の空文字が表示される

	//返り値を二つ設定することができる。値の取り出しに成功したならokはtrueと表示される
	s, ok := m4[3]
	if !ok {
		fmt.Println("error")
	}

	fmt.Println(s, ok)
	//もしsを_にすればvalueを受け取らないこともできる

	m4[2] = "USA"
	m4[3] = "CHINA"

	//値の取り消し
	delete(m4, 3)
	//第一引数にマップの変数名、第二引数には番号
	fmt.Println(m4)

	//要素数のlen関数
	fmt.Println(len(m4))

}

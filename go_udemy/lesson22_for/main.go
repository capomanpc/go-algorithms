package main

import "fmt"

//for

func main() {
	/*
		i := 0
		for {
			i++
			if i == 3 {
				break
			}
			fmt.Println("LOOP")
		}
	*/

	/*
		whire文

		point := 0
		for point < 10 {
			fmt.Println(point)
			point++
		}
	*/

	/*
		古典的for文

		for i := 0; i < 10; i++ {
			if i == 3 {
				continue
				//for文の処理に戻るのでこれより下は実行されない
			}
			fmt.Println(i)
		}
	*/

	/*
		配列の要素回繰り返すfor文

		arr := [3]int{1, 2, 3}
		for i := 0; i < len(arr); i++ {
			fmt.Println(arr[i])
		}
		len()は要素数を求める関数
	*/

	/*
		range...構文で使われるgolangの組み込みキーワード
				配列、スライス、文字列、マップ等に使用可能
				要素の数だけ繰り返し、反復処理の各ステップでindex番号と要素を取得できる

		arr := [3]int{1, 2, 3}

		for i, v := range arr {
			fmt.Println(i, v)
		}

		i...index番号
		v...配列の中身
		変数iのところを'_'とすることで破棄することができる
		変数vを使わない場合はfor i :=range arrと記述する
	*/

	/*
			スライス型のrangeを用いたfor文

			sl := []string{"Python", "PHP", "GO"}
			for i, v := range sl {
				fmt.Println(i, v)
		}
	*/

	//map-マップ
	m := map[string]int{"apple": 100, "banana": 200}

	for k, v := range m {
		fmt.Println(k, v)
	}

}

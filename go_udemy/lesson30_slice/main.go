package main

import "fmt"

//スライス
//宣言、操作

func main() {

	var sl []int

	//スライスの宣言
	//配列は要素数を指定してスライスは要素数を指定しない

	fmt.Println(sl)
	//[]だけが表示される

	//明示的な宣言、値も指定
	var sl2 []int = []int{100, 200}
	fmt.Println(sl2)

	//暗黙的な宣言
	sl3 := []string{"A", "B"}
	fmt.Println(sl3)

	//初期値は0
	sl4 := make([]int, 5)
	fmt.Println(sl4)

	//値の代入
	sl2[0] = 1000
	fmt.Println(sl2)

	//値の取り出し
	sl5 := []int{1, 2, 3, 4, 5}
	fmt.Println(sl5)
	fmt.Println(sl5[0])

	//2番から4番の手前までが表示される
	fmt.Println(sl5[2:4])

	//2番の手前まで表示される
	fmt.Println(sl5[:2])

	//2番から最後まで表示される
	fmt.Println(sl5[2:])

	//全部表示
	fmt.Println(sl5[:])

	//最後だけ表示-len関数を利用
	fmt.Println(sl5[len(sl5)-1])

	//最初と最後以外すべて取り出す
	fmt.Println(sl5[1 : len(sl5)-1])

}

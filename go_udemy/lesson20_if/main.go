package main

import "fmt"

//if
//条件分岐

func main() {
	a := 1
	if a == 2 {
		fmt.Println("two")
	} else if a == 1 {
		fmt.Println("one")
	} else {
		fmt.Println("I don't know")
	}

	//簡易文付きif文...条件式の前に簡易文を追加できる

	if b := 100; b == 100 {
		fmt.Println("one hundred")
	}

	x := 0
	if x := 2; true {
		fmt.Println(x)
	}
	fmt.Println(x)

	//if文で定義したx:=2はif文のスコープ内でのみ有効なので注意

}

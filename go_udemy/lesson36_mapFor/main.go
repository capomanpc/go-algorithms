package main

import "fmt"

//map
//for

func main() {
	m := map[string]int{
		"Apple":  100,
		"Banana": 200,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
	//kをアンダーバーにすることでvalueだけを取り出せる
	//valueを書かないことでkeyだけを取り出せる
	
}

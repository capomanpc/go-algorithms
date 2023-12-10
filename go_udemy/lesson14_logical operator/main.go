package main

import "fmt"

//論理演算子-logical operator-

func main() {
	fmt.Println(true && false == true)
	fmt.Println(true && true == true)
	fmt.Println(true && false == false)

	fmt.Println(true || false == false)
	fmt.Println(false || false == false)

	fmt.Println(!true)
	fmt.Println(!false)
	//論理値の反転	!論理値

}

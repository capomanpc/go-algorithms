package main

import "fmt"

//switch
//型switch

func anything(a interface{}) {
	fmt.Println(a)
	switch v := a.(type) {
	case int:
		fmt.Println(v + 10000)
	case string:
		fmt.Println(v + "!?")
	default:
		fmt.Println("I don't know")
	}

}

func main() {
	anything("aaa")
	anything(1)

	//interface{}型の変数xを型アサーションしている
	//アサーション-assertion-主張
	//型アサーションすることで演算や文字列の連結が可能になる

	var x interface{} = 3
	i := x.(int)
	//interface{}型のxを.(int)で型アサーションしてiに代入
	fmt.Println(i + 2)
	//fmt.Println(x + 2)

	f, isFloat64 := x.(float64)
	fmt.Println(f, isFloat64)

	//型アサーションが成功したらisFloat64にはTrueが代入される
	//isFloat64...変数xがfloat64型に変換可能かどうかを示すブール値

	//f:=x.(float64)はinterface{}型のxの中身の型と一致しないとfに代入されない

	fmt.Printf("isfloat64=%T\n", isFloat64)

	if x == nil {
		fmt.Println("None")
	} else if i, isInt := x.(int); isInt {
		fmt.Println(i, "x is Int")
	} else if s, isString := x.(string); isString {
		fmt.Println(s, isString)
	} else {
		fmt.Println("I don't know")
	}
	//isIntとisStringはbool型(True or false)なので実質的に条件式の役割を果たす
	//nilは中身がゼロという意味

	//型スイッチ
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("I don't know")
	}

	//x.(type)はswitch文に使われる特殊な構文(他にも存在するかも)

	switch v := x.(type) {
	case bool:
		fmt.Println(v, "bool")
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "string")
	}

	/*

		型スイッチのメリット

		f := x.(float64)
		このような処理があったときxの中身の型とfloat64が一致しなければ処理は行われない
		fにはinterface{}型の中身をfloat64型に変換したものが代入される

		xに代入された値の型が分からなければf := x.(float64)の処理を何度も行う必要がある
		そのときのために型スイッチが存在する


	*/
}

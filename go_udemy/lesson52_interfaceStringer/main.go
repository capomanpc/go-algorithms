package main

import "fmt"

//interface
//fmt.Stringer
//fmtパッケージに定義されているStringer型(interfaceで実装されている)

/*
type Stringer interface {
	String() string
}
*/

type Point struct {
	A int
	B string
}

func (p *Point) String() string {
	return fmt.Sprintf("<<%v, %v>>", p.A, p.B)
}

/*
StringメソッドはStringerインターフェースでまとめられている
fmt.Println関数で文字列生成をするときにStringerインターフェースを
満たすかチェックされている
満たしているときは文字列に対してStringメソッドが実行される

*/

func main() {
	p := &Point{100, "ABC"}
	fmt.Println(p)
	//fmt.Println関数が
}

package main

import "fmt"

//interface
//カスタムエラー

//interfaceを使って独自のエラーをカスタムエラーとして作る

/*
type error interface{
	Error() string
}

error型はinterfaceで上記のように定義されている
これが大元である
*/

type MyError struct {
	Message string
	ErrCode int
}

//MyError構造体はカスタムエラー情報を管理するために宣言

func (e *MyError) Error() string {
	return e.Message
}

//自分でエラーを作りたかったらMyErrorのErrorメソッドを作ればよい
//Errorメソッドはinterfaceでerror型になるように設定されている
//MyError型構造体と紐づいたError()メソッド
//エラーメッセージを返す

func RaiseError() error {
	return &MyError{Message: "カスタムエラーが発生しました", ErrCode: 1234}
}

//エラーメッセージとエラーコードの構造体を返す関数

func main() {
	err := RaiseError()
	fmt.Println(err.Error())
	/*
		//変数errにRaiseErrorで生成した構造体のアドレスを代入
		//RaiseError関数で得た構造体を使用してError関数の結果を出力

		//fmt.Println(err.Message)
		//上記のコードのerr構造体はinterface型なのでMessageフィールドを持たない
		//Go言語の error インターフェースは、単純なエラーメッセージ文字列を返す
		// Error() メソッドを持つことが規定されていますが、
		//Message フィールドそのものは error インターフェースには含まれていません。
		//そのため型アサーションをしてinterface型から元のMyError型に戻す
		//以下で型アサーションをすることでアクセスできるようにする
	*/

	e, ok := err.(*MyError)
	if ok {
		fmt.Println(e.ErrCode)
	}

	//型アサーション

	/*

		型アサーション	interface型の変数を具体的な型に変換するための操作

		value, ok := x.(Type)

		value	変数xをTypeに変換したものがvalueに代入される
		ok		bool型の変数、valueに代入出来たらTrue
		x		型アサーションしたい変数x
		Type	変換したいTypeのこと


	*/

}

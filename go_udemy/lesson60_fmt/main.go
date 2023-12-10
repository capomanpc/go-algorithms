package main

import "fmt"

//fmt

func main() {
	fmt.Println("表示")

	//fmt標準
	fmt.Print("Hello")

	//改行
	fmt.Println("Hello!")

	/*
		//Println系　各々の文字列は半角スペースで区切られ、文字列の最後に改行を追加
		fmt.Println("Hello","World!")
		fmt.Println("Hello","world!")
	*/

	/*
		//Printf系　フォーマットを指定
		fmt.Printf("%s\n","Hello")
		fmt.Printf("%#v\n","Hello")　//%#v 値の文法をそのまま出力する
	*/

	/*
		//Sprint系　出力ではなくフォーマットした結果を文字列で返す。変数に代入される。
		//String Printの略で返される結果が文字列であることを示している
		s := fmt.Sprint("Hello")
		s1 := fmt.Sprintf("%v\n","Hello")
		s2 := fmt.Sprintln("Hello")

		fmt.Println(s)
		fmt.Println(s1)
		fmt.Println(s2)

	*/

	/*
		//Fprint系　書き込み先指定
		//Fはファイルのこと
		fmt.Fprint(os.Stdout,"Hello")
		fmt.Fprintf(os.Stdout,"Hello")
		fmt.Fprintln(os.Stdout,"Hello")
		//Stdout:standard output"標準出力"

		f,_ := os.Create("test.txt")
		defer f.Close()

		fmt.Fprintln(f,"Fprint")
		//テキストファイルのfにFprintと入力している
		//func Fprintf(w io.Writer, format string, a ...any)
		//出力先は io.Writer型
	*/

}

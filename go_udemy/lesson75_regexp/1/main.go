package main

import (
	"fmt"
	"regexp"
)

//regexp-レジェックス、リージェックス
//Regular Expression 正規表現
//seikihyoug
//正規表現は文字列の表記方法のことであり、正規表現などで表される文字列を探すことをパターンマッチという
//正規表現という言葉は1940年代に日本の数学者が名付けた

func main() {

	//Goの正規表現の基本 regexp.MatchString 簡易的ver
	//regexp.MatchString(正規表現のパターン,検索対象の文字列)
	//返り値はbool値でmatchしたかチェック
	match, _ := regexp.MatchString("A", "ABC") //Aを含んでいるのでtrueを返す
	fmt.Println(match)                         //true

	//Compile
	//regexp.MatchString()は実行される度に正規表現のコンパイルをしているので非効率
	//同じ正規表現パターンを繰り返し使うのならば一度コンパイルをしてからそれを繰り返して使うべき
	//正規表現をコンパイルする関数
	re1, _ := regexp.Compile("A")  //コンパイルの結果をre1に代入
	match = re1.MatchString("ABC") //コンパイルしたre1で文字列を検索
	fmt.Println(match)

	//MustCompile
	//コンパイルに失敗した場合パニックを引き起こして終了する
	//誤ったパターンにコンパイルしてしまう事はその後のプログラムに大きな影響を及ぼすため強制終了になる
	//正常にコンパイルできていることを保証する
	re2 := regexp.MustCompile("A")
	match = re2.MatchString("ABC")
	fmt.Println(match)

	//regexp.MustCompile("\\d")　\dをエスケープしていて\d自体を検索する
	//regexp.MustCompile(`\d`)	\dという正規表現で特別な意味を持つもの

	//エスケープ...特定の文字が持つ特別な意味を無効化し、文字自体を表現すること
	//たとえばバックスラッシュ"\"はエスケープシーケンス(\nなど)の始まりを表し特別な意味を持つ
	//"\"この文字自体を使いたい場合"\\"のようにバックスラッシュを重ねることでエスケープしている
	//エスケープシーケンス..."\n"などの"\"で始まる特別な意味を持つ文字列
	re3 := regexp.MustCompile(`(?i)abc`)
	match = re3.MatchString("ABC")
	fmt.Println(match)

	re4 := regexp.MustCompile(`^ABC$`)
	match = re4.MatchString("ABC")
	fmt.Println(match)
	match = re4.MatchString("  ABC  ")
	fmt.Println(match)

	/*
		re := regexp.MustCompile("ab")
		re.MatchStrings("abc")
		//true
	*/

	re5 := regexp.MustCompile(".")
	match = re5.MatchString("ABC")
	fmt.Println(match)
	match = re5.MatchString("\n")
	fmt.Println(match)

	re6 := regexp.MustCompile("a+b*")
	fmt.Println(re6.MatchString("ab"))
	fmt.Println(re6.MatchString("a"))
	fmt.Println(re6.MatchString("aaaabbbbbbbb"))
	fmt.Println(re6.MatchString("b"))

	re7 := regexp.MustCompile("A+?A+?X")
	fmt.Println(re7.MatchString("AAX"))
	fmt.Println(re7.MatchString("AAAAAAXXXX"))

	re8 := regexp.MustCompile(`[XYZ]`)
	fmt.Println(re8.MatchString("Y"))

	re9 := regexp.MustCompile(`^[0-9A-Za-z_]{3}$`)
	fmt.Println(re9.MatchString("ABC"))
	fmt.Println(re9.MatchString("abcdefg"))

	re10 := regexp.MustCompile(`[^0-9A-Za-z_]`)
	fmt.Println(re10.MatchString("ABC"))
	fmt.Println(re10.MatchString("あ"))

	re1 = regexp.MustCompile(`(abc|ABC)(xyz|XYZ)`)
	fmt.Println(re1.MatchString("abcxyz"))
	fmt.Println(re1.MatchString("ABCXYZ"))
	fmt.Println(re1.MatchString("ABCxyz"))
	fmt.Println(re1.MatchString("ABCabc"))

	fs1 := re1.FindString("AAAAABCXYZZZZ")
	fmt.Println(fs1)
	fs2 := re1.FindAllString("ABCXYZABCXYZ", -1)
	fmt.Println(fs2)

	fmt.Println(re1.Split("ASHVJV<HABCXYZKNJBJVKABCXYZ", -1))
	re1 = regexp.MustCompile(`\s+`)
	fmt.Println(re1.Split("aaaaaaaaaa     bbbbbbb  cccccc", -1))

	fmt.Println(re1.ReplaceAllString("AAA BBB CCC", ","))
	re1 = regexp.MustCompile(`、|。`)
	fmt.Println(re1.ReplaceAllString("私は、Golangを使用する、プログラマー。", ""))

	re1 = regexp.MustCompile(`(\d+)-(\d+)-(\d+)`)
	s := `
	0123-456-7889
	111-222-333
	556-787-899
	`

	ms := re1.FindAllStringSubmatch(s, -1)
	for _, v := range ms {
		fmt.Println(v)
	}

	fmt.Println(re1.ReplaceAllString("Tel: 000-111-222", "$3-$2-$1"))
	fmt.Println(re1.ReplaceAllString("Tel: 000-111-222", "あ-い-う"))

}

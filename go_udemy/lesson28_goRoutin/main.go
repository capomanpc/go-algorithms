package main

import (
	"fmt"
	"time"
)

//go goroutin
//並行処理の一種 複数の処理を一つのプロセッサーで切り替えながら行う
//go文で簡単に並行処理が実行できる

/*
Goroutineとはスレッドより軽量な
*/

func sub() {
	for {
		fmt.Println("sub loop")
		time.Sleep(100 * time.Millisecond)
	}
}

//並行で動く関数sub

func main() {
	go sub()

	//goをつけるだけで並行処理を実現

	for {
		fmt.Println("Main Loop")
		time.Sleep(200 * time.Millisecond)
	}

	//main関数の中にこのfor文がないと正確に動作にしない
	/*
		下記のようなコードを考える

		func main(){
			go sub1()
			go sub2()
		}

		main関数が実行されるとgoroutineのsub1とsub2を呼び出した時点で処理が終了してしまう
		つまりsub1とsub2の中の処理の終了を待たずに全体の処理が終了してしまう
		main関数が終了するとき全体の処理が終了する
	*/
}

/*
	以下のように動作する

	Main Loop
	sub loop
	sub loop
	Main Loop
	sub loop
	sub loop
	Main Loop
	sub loop
	sub loop
	Main Loop
	sub loop
	sub loop
*/

/*
goroutineとスレッドは、両方ともプログラムの並行処理を実現するための仕組みですが、いくつかの違いがあります。

まず、goroutineはGo言語における独自の機能であり、スレッドはOSが提供する一般的な機能です。goroutineはGoラ
ンタイムによって管理され、複数のgoroutineを一つのスレッドにマルチプレクシングすることで、効率的に並行処理
を行います。スレッドはOSのカーネルによって管理され、プロセス内に複数のスレッドを作成することで、並行処理
を行います。

次に、goroutineはスタックサイズが少なく済むため、メモリ消費量が少ないだけでなく、スタックサイズを柔軟に
変更することも可能です1。スレッドはデフォルトで1MBや2MBのスタックサイズを持ち、メモリ消費量が多くなりま
す。また、スタックサイズを変更するにはOSの設定やコンパイラのオプションなどを変更する必要があります。

さらに、goroutineは協調的なスケジューリング方式を採用しており、goroutine自身が実行を中断したり再開した
りすることで、他のgoroutineにCPUを譲ります。スレッドはプリエンプティブなスケジューリング方式を採用して
おり、OSが定期的にスレッドの実行を中断したり再開したりすることで、他のスレッドにCPUを譲ります。

この違いにより、goroutineはコンテキストスイッチのコストが低くなります。コンテキストスイッチとは、実行
中の処理から別の処理に切り替える際に発生するオーバーヘッドのことです。goroutineは協調的な方式なので、
必要最小限の情報しか保存したり復元したりしません2。スレッドはプリエンプティブな方式なので、多くの情報
を保存したり復元したりします2。

以上のように、goroutineとスレッドは並行処理の仕組みとして似ていますが、実装や挙動に違いがあります。
goroutineはメモリ効率やパフォーマンスが高く、Go言語では簡単に使うことができます。しかし、goroutine
はGo言語固有の機能なので、他の言語と連携する際には注意が必要です。また、goroutine間の同期やデータ共
有も工夫が必要です。スレッドは汎用性や互換性が高く、多くの言語やOSで使うことができます。しかし、ス
レッドはメモリ効率やパフォーマンスが低く、作成や管理も複雑です。また、スレッド間の同期やデータ共有
も注意深く行う必要があります。
*/
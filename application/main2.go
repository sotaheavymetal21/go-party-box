package main

import "fmt"

// init 関数はプログラムの開始時に1回だけ実行されます。
func init() {
	fmt.Println("Initializing...")
}

// bazz関数に名前を渡して挨拶メッセージを出力するようにします。
func bazz(name string) {
	fmt.Printf("Hello, %s! Welcome to Bazz!\n", name)
}

func main() {
	bazz("Alice")
	fmt.Println("This is the main function!")
	bazz("Bob")
}

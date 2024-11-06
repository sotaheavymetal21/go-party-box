// package main

// import (
// 	"fmt"
// 	"os/user"
// 	"time"
// )

// func main() {
// 	fmt.Println("hello world", time.Now())
// 	fmt.Println(user.Current())
// }

// package main

// import "fmt"

// func main() {
// 	var i int = 1
// 	var f64 float64 = 1.2
// 	var s string = "test"
// 	var t, f bool = true, false
// 	fmt.Println(i, f64, s, t, f)

// 	xi := 1
// 	xf64 := 1.2
// 	xs := "test"
// 	xt, xf := true, false
// 	fmt.Println(xi, xf64, xs, xt, xf)
// }

// package main

// import "fmt"

// const Pi = 3.14

// const (
// 	Username = "test_user"
// 	Password = "test_pass"
// )

// var big int = 922372036854775807 + 2

// func main() {
// 	fmt.Println(Pi, Username, Password)
// 	fmt.Println(big - 1)
// }

// package main

// import "fmt"

// func main() {
// 	var (
// 		u8  uint8     = 255
// 		i8  int8      = 127
// 		f32 float32   = 0.2
// 		c64 complex64 = -5 + 12i
// 	)

// 	fmt.Println(u8, i8, f32, c64)
// 	fmt.Printf("type=%T value=%v", u8, u8)

// }

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello World")            // "Hello World" をそのまま出力
	fmt.Println("Hello " + "World")       // 文字列の結合を出力
	fmt.Println(string("Hello World"[0])) // "Hello World" の先頭文字 'H' を出力

	var s string = "Hello World" // 変数 s に "Hello World" を代入

	s = strings.Replace(s, "H", "X", 1) // s 内の "H" を "X" に置き換える（最初の 1 つだけ）
	fmt.Println(s)                      // 置き換え後の文字列 "Xello World" を出力

	fmt.Println(strings.Contains(s, "World")) // s に "World" が含まれているかをチェックし、結果 (true) を出力

	fmt.Println(`Test
		Test`) // バッククォートで囲んだ文字列をそのまま出力（改行含む）

	fmt.Println("\"") // ダブルクォートを出力（エスケープシーケンスを使用）
	fmt.Println(`"`)  // バッククォートで囲んだダブルクォートを出力
}

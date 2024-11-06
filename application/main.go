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

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	fmt.Println("Hello World")            // "Hello World" をそのまま出力
// 	fmt.Println("Hello " + "World")       // 文字列の結合を出力
// 	fmt.Println(string("Hello World"[0])) // "Hello World" の先頭文字 'H' を出力

// 	var s string = "Hello World" // 変数 s に "Hello World" を代入

// 	s = strings.Replace(s, "H", "X", 1) // s 内の "H" を "X" に置き換える（最初の 1 つだけ）
// 	fmt.Println(s)                      // 置き換え後の文字列 "Xello World" を出力

// 	fmt.Println(strings.Contains(s, "World")) // s に "World" が含まれているかをチェックし、結果 (true) を出力

// 	fmt.Println(`Test
// 		Test`) // バッククォートで囲んだ文字列をそのまま出力（改行含む）

// 	fmt.Println("\"") // ダブルクォートを出力（エスケープシーケンスを使用）
// 	fmt.Println(`"`)  // バッククォートで囲んだダブルクォートを出力
// }

// package main

// import "fmt"

// func main() {
// 	// 変数 t と f をそれぞれ bool 型で初期化
// 	var t, f bool = true, false

// 	// 変数の型(%T)と値(%v)、論理値(%t)を出力
// 	// t の場合: 型は "bool"、値は "1"、論理値は "true"
// 	fmt.Printf("%T %v %t\n", t, 1, t)

// 	// f の場合: 型は "bool"、値は "0"、論理値は "false"
// 	fmt.Printf("%T %v %t\n", f, 0, f)

// 	// 論理演算子の AND 演算子 "&&" の結果を出力
// 	// 両方が true の場合のみ true になる
// 	fmt.Println(true && true)   // true
// 	fmt.Println(true && false)  // false
// 	fmt.Println(false && false) // false

// 	// 論理演算子の OR 演算子 "||" の結果を出力
// 	// 少なくとも一方が true なら true になる
// 	fmt.Println(true || true)   // true
// 	fmt.Println(true || false)  // true
// 	fmt.Println(false || false) // false

// 	// 論理演算子の NOT 演算子 "!" の結果を出力
// 	// true を false に、false を true に反転させる
// 	fmt.Println(!true)  // false
// 	fmt.Println(!false) // true
// }

package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 整数 x を定義して float64 に型変換し、変数 xx に代入
	var x int = 1
	xx := float64(x)
	// xx の型(%T)、値(%v)、浮動小数点形式(%f)で出力
	fmt.Printf("%T %v %f\n", xx, xx, xx)

	// 浮動小数点数 y を定義して int に型変換し、変数 yy に代入
	var y float64 = 1.2
	yy := int(y)
	// yy の型(%T)、値(%v)、整数形式(%d)で出力
	fmt.Printf("%T %v %d\n", yy, yy, yy)

	// 文字列 s を整数に変換し、変数 i に代入
	var s string = "14"
	i, _ := strconv.Atoi(s)
	// i の型(%T)、値(%v)で出力
	fmt.Printf("%T %v\n", i, i)

	// 文字列 h の先頭文字を取得し、文字として出力
	h := "Hello World"
	fmt.Println(string(h[0]))
}

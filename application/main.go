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

package main

import "fmt"

const Pi = 3.14

const (
	Username = "test_user"
	Password = "test_pass"
)

var big int = 922372036854775807 + 2

func main() {
	fmt.Println(Pi, Username, Password)
	fmt.Println(big - 1)
}

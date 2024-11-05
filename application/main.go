package main

import (
	"fmt"
	"os/user"
	"time"
)

func main() {
	fmt.Println("hello world", time.Now())
	fmt.Println(user.Current())
}

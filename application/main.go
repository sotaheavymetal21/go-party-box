package main

import (
	"fmt"
	"net/http"

	"github.com/sotaheavymetal21/go-party-box/routers"
)

func main() {
	// ルーターを初期化
	r := routers.NewRouter()

	// サーバーをポート 8000 で起動
	fmt.Println("Starting server on port 8000...")
	http.ListenAndServe(":8000", r)
}

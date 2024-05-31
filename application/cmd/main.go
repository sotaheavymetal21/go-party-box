// main.go

// メインパッケージの宣言
package main

// 必要なパッケージのインポート
import (
	"log"
	"net/http"

	"github.com/sotaheavymetal21/go-party-box/application/handlers" // 自作パッケージのインポート
)

// main 関数の定義
func main() {
	// "/weather" エンドポイントに WeatherHandler 関数を割り当てる
	http.HandleFunc("/weather", handlers.WeatherHandler)

	// サーバーをポート 8080 で起動し、エラーが発生した場合はログに出力する
	log.Fatal(http.ListenAndServe(":8080", nil))
}

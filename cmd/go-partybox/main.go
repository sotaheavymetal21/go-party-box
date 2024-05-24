package main

import (
	"net/http"
	"path/filepath"
	"runtime"
)

func main() {
	// カレントディレクトリを取得
	_, filename, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(filename)

	// handlersパッケージの相対パス
	handlersDir := filepath.Join(currentDir, "../handlers")

	// ハンドラのインポート
	http.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		// この関数内でのハンドラの処理
	})

	http.ListenAndServe(":8080", nil)
}

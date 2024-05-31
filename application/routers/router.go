package routers

import (
	"encoding/json"
	"net/http"
)

// HealthResponse 構造体はヘルスチェックAPIのレスポンスを表します
type HealthResponse struct {
	Status string `json:"status"`
}

// /health エンドポイントのハンドラ関数
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// HealthResponse構造体を作成
	response := HealthResponse{
		Status: "ok",
	}

	// JSON形式にエンコードしてレスポンスを返す
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

package handlers

import (
	"fmt"
	"net/http"
)

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	apiKey := "YOUR_API_KEY"
	city := r.URL.Query().Get("city")
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to get weather data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// レスポンスの読み込みなど、天気情報の取得処理を記述します

	// レスポンスをクライアントに返す
	// w.Write(responseData)
}

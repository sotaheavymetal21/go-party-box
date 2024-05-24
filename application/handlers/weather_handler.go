// handlers/weather_handler.go

// handlers パッケージの宣言
package handlers

// 必要なパッケージのインポート
import (
	"encoding/json"
	"fmt"
	"net/http"
)

// WeatherResponse 構造体は、OpenWeatherMap APIからの天気情報のレスポンスを表します。
type WeatherResponse struct {
	City    string `json:"name"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

// WeatherHandler 関数は、"/weather" エンドポイントに対するハンドラです。
// クエリパラメータとして指定された都市の天気情報を取得し、レスポンスとして返します。
func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	// URLからクエリパラメータとして都市を取得
	city := r.URL.Query().Get("city")
	// OpenWeatherMap APIのAPIキー
	apiKey := "YOUR_API_KEY"
	// APIエンドポイントのURLを構築
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)

	// 天気情報APIにリクエストを送信し、レスポンスを取得
	resp, err := http.Get(url)
	if err != nil {
		// エラーが発生した場合は、HTTPエラーを返して終了
		http.Error(w, "Failed to get weather data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// レスポンスのJSONをWeatherResponse構造体にデコード
	var weatherData WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
		// デコードに失敗した場合は、HTTPエラーを返して終了
		http.Error(w, "Failed to decode weather data", http.StatusInternalServerError)
		return
	}

	// レスポンスを整形してクライアントに返す
	response := fmt.Sprintf("Weather in %s: %s", weatherData.City, weatherData.Weather[0].Description)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

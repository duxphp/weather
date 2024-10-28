package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/cast"
	"io"
	"net/http"
)

var accountKey = "P3JOHAV73NJQQjRCu"
var accountSecret = "S-ktM5BwJbNCX_CRa"
var apiUri = "https://api.seniverse.com/v3/weather/now.json"

func main() {

	router := mux.NewRouter()

	// 渲染编译后的首页
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			_, _ = w.Write([]byte("Hello, This is GET method. [Go]"))
		case http.MethodPost:
			_, _ = w.Write([]byte("Hello, This is POST method. [Go]"))
		}
	})

	// 天气查询api
	router.HandleFunc("/api/weather/{city}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		city := vars["city"]
		resp, err := http.Get(apiUri + "?key=" + accountSecret + "&location=" + city + "&language=zh-Hans&unit=c")
		if err != nil {
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			RespJson(w, &RespData{
				Status:  500,
				Message: "请求接口失败",
				Data:    map[string]any{},
			})
			return
		}

		data := map[string]any{}
		if err := json.Unmarshal(body, &data); err != nil {
			RespJson(w, &RespData{
				Status:  500,
				Message: "接口数据解析失败",
				Data:    map[string]any{},
			})
			return
		}

		if data["results"] == nil {
			RespJson(w, &RespData{
				Status:  500,
				Message: cast.ToString(data["status"]),
				Data:    map[string]any{},
			})
			return
		}

		RespJson(w, &RespData{
			Status:  200,
			Message: "请求成功",
			Data:    data["results"],
		})

	})

	fmt.Println("Server is running at http://0.0.0.0:8800")
	_ = http.ListenAndServe(":8800", router)
}

type RespData struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func RespJson(w http.ResponseWriter, resp *RespData) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

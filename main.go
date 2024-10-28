package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/cast"
	"io"
	"io/fs"
	"net/http"
)

var accountSecret = "S-ktM5BwJbNCX_CRa"
var apiUri = "https://api.seniverse.com/v3/weather/now.json"

//go:embed dist
var distDir embed.FS

func main() {

	router := mux.NewRouter()

	// 天气查询路由
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

		// 处理异常
		if data["results"] == nil {
			RespJson(w, &RespData{
				Status:  500,
				Message: cast.ToString(data["status"]),
				Data:    map[string]any{},
			})
			return
		}

		// 返回结果
		RespJson(w, &RespData{
			Status:  200,
			Message: "请求成功",
			Data:    data["results"],
		})

	})

	// 静态文件路由
	dist, err := fs.Sub(distDir, "dist")
	if err != nil {
		panic(err)
	}
	dir := http.FileServer(http.FS(dist))
	router.PathPrefix("/").Handler(dir)

	fmt.Println("Server is running at http://0.0.0.0:8800")
	_ = http.ListenAndServe(":8800", router)
}

// RespData 返回结构
type RespData struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// RespJson 返回json封装
func RespJson(w http.ResponseWriter, resp *RespData) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 上传目录

func UploadVideo(w http.ResponseWriter, r *http.Request) {
	// 设置响应头为 JSON 格式
	w.Header().Set("Content-Type", "application/json")

	// 返回成功响应
	response := map[string]string{"status": "success"}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
	}
}

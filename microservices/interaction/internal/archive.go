package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	auth "github.com/Yux77Yux/platform_backend/pkg/auth"
)

func Archive(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPatch:
		SaveArchive(w, r)
	case http.MethodPost:
		UploadArchive(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
func SaveArchive(w http.ResponseWriter, r *http.Request) {
	type RequestBody struct {
		Order       int `json:"order"`
		AccessToken struct {
			Value string `json:"value"`
		} `json:"accessToken"`
	}
	var body RequestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	token := body.AccessToken.Value
	pass, userId, err := auth.Auth("get", "interaction", token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"err": err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	if !pass {
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"err": err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// 获取存档的行为数据
	order := body.Order
	if order < 0 || order > 3 {
		http.Error(w, "Failed to upload file", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()
	tmpFile, err := cache.GetArchive(ctx, userId, strconv.Itoa(order))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"err": err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", tmpFile.Name()))

	http.ServeContent(w, r, tmpFile.Name(), time.Now(), tmpFile)
}
func UploadArchive(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(1 << 22)
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	accessTokenStr := r.FormValue("accessToken")

	type TokenWrapper struct {
		Value string `json:"value"`
	}

	var tokenWrapper TokenWrapper
	err = json.Unmarshal([]byte(accessTokenStr), &tokenWrapper)
	if err != nil {
		http.Error(w, "Invalid accessToken JSON", http.StatusBadRequest)
		return
	}

	token := tokenWrapper.Value
	pass, userId, err := auth.Auth("post", "interaction", token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"err": err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	if !pass {
		response := map[string]string{"err": "not pass"}
		json.NewEncoder(w).Encode(response)
		return
	}

	order := r.FormValue("order")
	file, _, err := r.FormFile("archive")
	if err != nil {
		http.Error(w, "Failed to upload file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	err = cache.SetArchive(ctx, userId, order, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"err": err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"success": "OK!"}
	json.NewEncoder(w).Encode(response)
}

func ArchiveOrder(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		SetArchiveOrder(w, r)
	case http.MethodGet:
		GetArchiveOrder(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
func SetArchiveOrder(w http.ResponseWriter, r *http.Request) {
	type RequestBody struct {
		Order       int `json:"order"`
		AccessToken struct {
			Value string `json:"value"`
		} `json:"accessToken"`
	}
	var body RequestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	token := body.AccessToken.Value
	pass, userId, err := auth.Auth("update", "interaction", token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"err": err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	if !pass {
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"err": err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// 获取存档的行为数据
	order := body.Order
	if order < 0 || order > 3 {
		http.Error(w, "Failed to upload file", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()
	err = cache.SetUsingArchive(ctx, userId, strconv.Itoa(order))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"err": err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	go func(userId int64, ctx context.Context) {
		err = messaging.SendMessage(ctx, EXCHANGE_COMPUTE_USER, KEY_COMPUTE_USER, &common.UserDefault{
			UserId: userId,
		})
	}(userId, ctx)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"success": "OK!"}
	json.NewEncoder(w).Encode(response)
}
func GetArchiveOrder(w http.ResponseWriter, r *http.Request) {
	userIdStr := r.URL.Query().Get("userId")
	id, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()
	order, exist, err := cache.GetUsingArchive(ctx, id)
	if err != nil {
		log.Printf("redis err: %v \n", err)
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"err": err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]any{"order": order, "exist": exist}
	json.NewEncoder(w).Encode(response)
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 定义 API 响应数据结构
type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

// 简单的 get 请求
func GetHandler(w http.ResponseWriter, r *http.Request) {
	// 创建返回数据
	response := Response{
		Message: "Hello, this is your API response!",
		Status:  200,
	}

	// 将数据转换为 JSON
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error generating JSON", http.StatusInternalServerError)
		return
	}

	// 发送 JSON 响应
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// 简单的 POST 请求
func PostHandler(w http.ResponseWriter, r *http.Request) {
	// 确保是 POST 请求
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// 解析 JSON 请求体
	type User struct {
		Username string
		Email    string
	}
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	// 创建返回数据
	response := Response{
		Message: fmt.Sprintf("User [%s] with email [%s] registered successfully!", user.Username, user.Email),
		Status:  201,
	}

	// 发送 JSON 响应
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

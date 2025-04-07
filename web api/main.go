package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	_ "github.com/mattn/go-sqlite3"
)

// 连接数据库
var db = InitDB()

// 处理 Token 刷新请求
func refreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Token string `json:"token"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// 解析 JWT
	token, err := ParseJWT(req.Token)
	if err != nil || !token.Valid {
		http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
		return
	}

	// 获取 JWT 的声明
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["username"] == nil {
		http.Error(w, "Invalid token claims", http.StatusUnauthorized)
		return
	}

	username := claims["username"].(string)
	_, role, _ := GetUser(db, username)

	// 生成新的 JWT
	newToken, err := GenerateJWT(username, role)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": newToken})
}

// 用户注册
func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// 默认角色为 "user"，仅允许 "admin" 或 "user"
	if req.Role != "admin" {
		req.Role = "user"
	}
	hashpwd, _ := HashPassword(req.Password)
	// 插入数据库
	_, err := db.Exec("INSERT INTO users (username, password, role) VALUES (?, ?, ?)", req.Username, hashpwd, req.Role)
	if err != nil {
		http.Error(w, "Username already exists", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

// 用户登录
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST is allowed", http.StatusMethodNotAllowed)
		return
	}

	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// 查询用户
	storedPassword, role, err := GetUser(db, credentials.Username)
	if err != nil || !checkPassword(storedPassword, credentials.Password) {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// 生成 JWT，包含角色信息
	token, err := GenerateJWT(credentials.Username, role)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// 受保护路由（必须携带 JWT）
func protectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "You have accessed a protected route!"})
}
func main() {
	InitDB()

	mux := http.NewServeMux()
	// 绑定路由
	mux.HandleFunc("/api/get", GetHandler)
	mux.HandleFunc("/api/post", PostHandler)
	mux.HandleFunc("/api/login", loginHandler) // 用户登录
	mux.HandleFunc("/api/register", registerHandler)
	mux.Handle("/api/protected", JwtMiddleware(http.HandlerFunc(protectedHandler))) // 受保护路由
	mux.HandleFunc("/api/refresh-token", refreshTokenHandler)                       // 刷新 Token

	// 使用 CORS 中间件包装 Mux
	handler := CorsMiddleware(mux)

	// 启动服务器
	fmt.Println("Server is running on port 8080...")
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

// [INFO]
/*
JWT 本身是 无状态的，不存储在服务器端，因此无法像 session 一样直接销毁。
常见的 JWT 失效方案：

黑名单（将已登出的 Token 存储，并检查是否无效）
Token 旋转（每次刷新 Token 都会让旧 Token 失效）
短生命周期 + 频繁刷新（不登出，Token 短时间内自动失效）
*/

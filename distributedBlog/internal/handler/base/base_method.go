package base

import (
	"crypto/sha256"
	"distributedBlog/internal/const_values"
	"encoding/hex"
	"net/http"
)

// 处理预检请求
func Prefix_Managing(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" || r.Method == "POST" {
		// 这是预检请求 ，在POST请求也对其进行处理
		(w).Header().Set("Access-Control-Allow-Origin", const_values.BACKEND_SOURCE) // 允许指定来源的跨域请求
		(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		(w).Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, X-Extra-Header, Content-Type, Accept, Authorization")
		(w).Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		(w).Header().Set("Access-Control-Allow-Credentials", "true")
	}

}

// 跨域请求处理方法
func CORS_Managing(w *http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	if origin == const_values.FRONTEND_DOMAIN_NAME {
		(*w).Header().Set("Access-Control-Allow-Origin", origin)
		(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		(*w).Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, X-Extra-Header, Content-Type, Accept, Authorization")
		(*w).Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	} else {
		(*w).WriteHeader(http.StatusForbidden)
		return
	}
}
func CORS_ALLOW_ALL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", const_values.FRONTEND_ALL_DOMAIN)
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, X-Extra-Header, Content-Type, Accept, Authorization")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}

// 对于数据进行编码操作
func Hash_encoder(in string) string {
	hash := sha256.Sum256([]byte(in))
	hashString := hex.EncodeToString(hash[:])
	return hashString
}

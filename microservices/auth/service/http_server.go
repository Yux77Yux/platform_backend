package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"

	client "github.com/Yux77Yux/platform_backend/microservices/auth/client"
)

func HttpServerRun() {
	mux := http.NewServeMux()
	UserHandlers(mux)

	wrap_handler := ApplyMiddlewares(mux, CorsMiddleware)

	if err := http.ListenAndServe(http_addr, wrap_handler); err != nil {
		log.Println("error: http server ", err)
	} else {
		log.Println("http server run in ", http_addr)
	}
}

type Middleware func(http.Handler) http.Handler

func ApplyMiddlewares(handler http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}

	return handler
}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 设置 CORS 头
		w.Header().Set("Access-Control-Allow-Credentials", "true")                                                                                                                            // 允许的来源
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")                                                                                                                // 允许的来源
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT,PATCH, DELETE, OPTIONS")                                                                                               // 允许的方法
		w.Header().Set("Access-Control-Allow-Headers", "DNT,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Range,Authorization,Access-Control-Allow-Origin,Cookie") // 允许的请求头

		// 处理预检请求
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

type AccessTokenType struct {
	Value     string                 `json:"value,omitempty"`
	ExpiresAt *timestamppb.Timestamp `json:"expires_at,omitempty"`
}

type MsgType struct {
	Status  string `json:"status,omitempty"`
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Details string `json:"details,omitempty"`
	TraceId string `json:"trace_id,omitempty"`
}

func refreshByHttp(w http.ResponseWriter, r *http.Request) {
	log.Println("cookie")
	// 提取 httpOnly cookie
	cookie, err := r.Cookie("refreshToken")
	if err != nil {
		log.Println(r)
		http.Error(w, "Cookie not found", http.StatusUnauthorized)
		return
	}
	log.Println(cookie)
	// 创建 gRPC 请求上下文，并将 cookie 作为元数据传递
	md := metadata.Pairs("cookie", cookie.Value)
	ctx := metadata.NewOutgoingContext(r.Context(), md)

	// 调用 gRPC 服务
	client, err := client.NewAuthClient()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to call gRPC service: %v", err), http.StatusInternalServerError)
		return
	}
	resp, err := client.Refresh(ctx)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to call gRPC service: %v", err), http.StatusInternalServerError)
		return
	}

	// 获取 gRPC 响应中的数据
	accessToken := resp.GetAccessToken() // 假设是获取 AccessToken 的方法
	apiResponse := resp.GetMsg()         // 获取 ApiResponse

	log.Println("get: ", accessToken)

	// 创建一个结构体来表示 HTTP 响应
	httpResponse := struct {
		AccessToken *AccessTokenType `json:"access_token"`
		Msg         *MsgType         `json:"msg"`
	}{
		AccessToken: &AccessTokenType{
			Value:     accessToken.GetValue(),
			ExpiresAt: accessToken.GetExpiresAt(),
		},
		Msg: &MsgType{
			Status:  apiResponse.GetStatus().String(),
			Code:    apiResponse.GetCode(),
			Message: apiResponse.GetMessage(),
			Details: apiResponse.GetDetails(),
			TraceId: apiResponse.GetTraceId(),
		},
	}

	// 设置响应头为 JSON 格式
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// 将响应写入 HTTP 响应体
	if err := json.NewEncoder(w).Encode(httpResponse); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
	}
}

func UserHandlers(mux *http.ServeMux) {
	router := GetRouter(mux)

	router.Handler("/api/http/auth/refresh", refreshByHttp)
}

type handlerFunc func(http.ResponseWriter, *http.Request)

type CustomRouter struct {
	mux *http.ServeMux
}

func (r *CustomRouter) Handler(api string, handler handlerFunc) {
	r.mux.Handle(api, http.HandlerFunc(handler))
}

func GetRouter(mux *http.ServeMux) CustomRouter {
	CustomRouter := &CustomRouter{
		mux: mux,
	}
	return *CustomRouter
}

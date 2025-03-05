package service

import (
	// "encoding/json"
	// "fmt"
	"log"
	"net/http"
	// "google.golang.org/grpc/metadata"
	// "google.golang.org/protobuf/types/known/timestamppb"
)

func Handlers(mux *http.ServeMux) {
	router := GetRouter(mux)

	router.Handler("/http/creation", UploadVideo)
}

func HttpServerRun() {
	mux := http.NewServeMux()
	Handlers(mux)

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

type handlerFunc = func(http.ResponseWriter, *http.Request)

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

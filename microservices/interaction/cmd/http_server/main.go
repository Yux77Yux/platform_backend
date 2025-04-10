package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	internal "github.com/Yux77Yux/platform_backend/microservices/interaction/internal"
	middlewares "github.com/Yux77Yux/platform_backend/pkg/middlewares"
)

func ServerArchiveMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/archive", internal.Archive)
	mux.HandleFunc("/api/archive/order", internal.ArchiveOrder)

	return mux
}

func main() {
	mux := ServerArchiveMux()
	wrapHandler := middlewares.ApplyMiddlewares(mux, middlewares.CorsMiddleware)

	srv := &http.Server{
		Addr:    ":50041",
		Handler: wrapHandler,
	}

	// 创建一个 channel 用于监听退出信号（如 Ctrl+C）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 开一个 goroutine 启动 HTTP 服务
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()
	log.Println("HTTP server started on :50041")

	// 等待退出信号
	<-quit
	log.Println("Shutting down HTTP server...")

	// 创建 context，设置最大关闭等待时间
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("HTTP server forced to shutdown: %v", err)
	}

	log.Println("HTTP server exited properly")
}

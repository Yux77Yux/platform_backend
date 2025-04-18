package HttpServer

import (
	"context"
	"log"
	"net/http"
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

func Run() func(chan any) {
	mux := ServerArchiveMux()
	wrapHandler := middlewares.ApplyMiddlewares(mux, middlewares.CorsMiddleware)

	srv := &http.Server{
		Addr:    ":50041",
		Handler: wrapHandler,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()
	log.Println("HTTP server started on :50041")

	return func(signal chan any) {
		log.Println("Shutting down HTTP server...")

		// 创建 context，设置最大关闭等待时间
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalf("HTTP server forced to shutdown: %v", err)
		}

		close(signal)
		log.Println("HTTP server exited properly")
	}
}

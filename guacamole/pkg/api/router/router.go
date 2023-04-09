package router

import (
	"net/http"

	"github.com/tatsuyayamauchi/golang-http-server-example/guacamole/pkg/api/handler"
	"github.com/tatsuyayamauchi/golang-http-server-example/guacamole/pkg/api/middleware"
)

func Initialize() http.Handler {
	mux := http.NewServeMux()
	middlewares := middleware.Middlewares{
		middleware.NewLoggingMiddleware().Middleware,
		middleware.NewAuthMiddleware().Middleware,
		middleware.NewCorsMiddleware().Middleware,
		middleware.NewRecoveryMiddleware().Middleware,
	}

	mux.HandleFunc("/healthcheck",
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		},
	)

	mux.HandleFunc("/echo", handler.NewEchoHandler().Echo)

	return middlewares.Chain(mux)
}

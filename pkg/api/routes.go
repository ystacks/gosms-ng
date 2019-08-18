/**
 * File              : routes.go
 * Author            : Jiang Yitao <jiangyt.cn#gmail.com>
 * Date              : 11.08.2019
 * Last Modified Date: 11.08.2019
 * Last Modified By  : Jiang Yitao <jiangyt.cn#gmail.com>
 */
package api

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
	. "github.com/ystacks/gosms-ng/logger"
	"github.com/ystacks/gosms-ng/pkg/api/config"
	"github.com/ystacks/gosms-ng/pkg/api/sms"
	"go.uber.org/zap"
)

func Routes() *chi.Mux {

	r := chi.NewMux()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.RedirectSlashes)
	//r.Use(utils.Logger(logger))
	r.Use(middleware.Recoverer)

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			//ctx := context.WithValue(req.Context(), utils.ContextDBName, tx)
			//next.ServeHTTP(w, req.WithContext(ctx))
		})
	})

	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/api/v1", func(r chi.Router) {
		r.Use(apiVersionCtx("v1"))
		r.Mount("/sms", sms.Routes())
		r.Mount("/admin", config.Routes())
	})

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		Logger.Info("route", zap.String("method", method), zap.String("route", route)) // Walk and print out all routes
		return nil
	}
	if err := chi.Walk(r, walkFunc); err != nil {
		Logger.Fatal("Logging err", zap.Error(err)) // panic if there is an error
	}
	return r
}

func apiVersionCtx(version string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(context.WithValue(r.Context(), "api.version", version))
			next.ServeHTTP(w, r)
		})
	}
}

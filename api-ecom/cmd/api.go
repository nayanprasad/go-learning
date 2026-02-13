package main

import (
	"api-ecom/internal/product"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type appplication struct {
	config config
	//logger
	//db driver
}

func (*appplication) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	productHanlder := product.NewHandler(nil)
	r.Get("/products", productHanlder.ListProducts)

	return r
}

func (app *appplication) run(h http.Handler) error {
	svr := &http.Server{
		Addr:    app.config.addr,
		Handler: h,
	}
	slog.Info("Server has started", "port", app.config.addr)
	return svr.ListenAndServe()
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}

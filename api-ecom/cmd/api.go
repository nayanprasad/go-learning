package main

import (
	"api-ecom/internal/product"
	"log/slog"
	"net/http"

	repo "api-ecom/internal/adapters/psql/sqlc"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
)

type appplication struct {
	config config
	//logger
	db *pgx.Conn
}

func (app *appplication) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	productService := product.NewService(repo.New(app.db))
	productHanlder := product.NewHandler(productService)
	r.Get("/products", productHanlder.ListProducts)
	r.Post("/product/create", productHanlder.CreateProduct)

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

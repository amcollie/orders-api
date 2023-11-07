package application

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/amcollie/orders-api/handler"
	"github.com/amcollie/orders-api/repository/order"
)

func (a *App) loadRoutes() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader((http.StatusOK))
	}).Methods("GET")

	s := router.PathPrefix("/orders").Subrouter()
	a.loadOrderRoutes(s)

	a.router = router
}

func (a *App) loadOrderRoutes(router *mux.Router) {
	orderHandler := &handler.Order{
		Repo: &order.RedisRepo{
			Client: a.rdb,
		},
	}

	router.HandleFunc("", orderHandler.Create).Methods("POST")
	router.HandleFunc("/", orderHandler.Create).Methods("POST")
	router.HandleFunc("", orderHandler.List).Methods("GET")
	router.HandleFunc("/", orderHandler.List).Methods("GET")
	router.HandleFunc("/{id}", orderHandler.GetByID).Methods("GET")
	router.HandleFunc("/{id}", orderHandler.UpdateByID).Methods("PUT")
	router.HandleFunc("/{id}", orderHandler.DeleteByID).Methods("DELETE")
}

package handler

import (
	"fmt"
	"net/http"
)

type Order struct{}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create an order!")
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "List all orders!")
}

func (o *Order) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get an order by id")
}

func (o *Order) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update an order by id")
}

func (o *Order) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete an order by id")
}

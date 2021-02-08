package app

import (
	"net/http"

	"github.com/Outerwolf/bookstore_items-api/controllers"
)

func mapUrls() {
	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)

	router.HandleFunc("/items", controllers.ItemsCoontroller.Create).Methods(http.MethodPost)
	router.HandleFunc("/items{id}", controllers.ItemsCoontroller.Get).Methods(http.MethodGet)

	router.HandleFunc("/items{id}", controllers.ItemsCoontroller.Search).Methods(http.MethodPost)
}

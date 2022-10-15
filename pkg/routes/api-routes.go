package routes

import (
	"github.com/gorilla/mux"
	controllers "github.com/timalbpathirana/golangAPI/pkg/controller"
)

var RegisterStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/products", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/addProduct", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/product/{id}", controllers.GetProductById).Methods("GET")
	router.HandleFunc("/product/{id}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/product/{id}", controllers.DeleteProductById).Methods("DELETE")
}

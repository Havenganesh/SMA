package main

import (
	"fmt"
	"net/http"
	"sma/db"
	rpc "sma/rpcGateway"
	"sma/service"
	"sma/validate"

	"github.com/gorilla/mux"
)

func main() {
	db.Init()
	rpc.Init()
	service.Init()
	validate.Init()
	appInit()
}

func appInit() {
	// Set up Gorilla Mux router
	router := mux.NewRouter()

	// Public routes
	router.HandleFunc("/api/{service}/{method}", rpc.RPCHandler).Methods("POST")

	protected := router.PathPrefix("/rpc").Subrouter()
	protected.Use(rpc.JwtMiddleware)
	protected.HandleFunc("/{service}/{method}", rpc.RPCHandler).Methods("POST")

	// Start the server
	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", router)
}

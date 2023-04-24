package main

import (
	"cep-consult/api/handle"
	_ "cep-consult/docs"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title API de consulta de endereço
// @description Esta é uma API que consulta informações de endereço a partir de um CEP.
// @version 1
// @host localhost:8080
// @BasePath /

func main() {
	r := mux.NewRouter()

	r.PathPrefix("/docs/").Handler(http.StripPrefix("/docs/", http.FileServer(http.Dir("./docs"))))
	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	r.HandleFunc("/v1/consult-address", handle.ConsultAdress).Methods("POST")

	http.ListenAndServe(":8080", r)
}

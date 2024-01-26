package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"fmt"
)

type Usuario struct {
	Nome string `json:"nome,omitempty"`
}

func GetUsuario(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

	usuario := Usuario{
		Nome: "David",
	}

    json.NewEncoder(w).Encode(usuario)
}

func main() {
    router := mux.NewRouter()

    // Endpoints
    router.HandleFunc("/usuario", GetUsuario).Methods("GET")

	fmt.Printf("Iniciando programa na porta 8080")
    http.ListenAndServe(":8080", router)
}
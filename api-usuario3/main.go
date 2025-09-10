package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Usuario struct {
	Nome string `json:"nome,omitempty"`
}

func GetUsuario(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	log.Print("[API 3] - Acessando dados de usuário")

	usuario := Usuario{
		Nome: "TigreDoMexico - API 3",
	}

	for key, values := range r.Header {
		for _, value := range values {
			log.Printf("[API 3] Header: %s = %s", key, value)
		}
	}

	if err := json.NewEncoder(w).Encode(usuario); err != nil {
		log.Printf("Erro ao codificar JSON: %v", err)
		http.Error(w, "Erro interno", http.StatusInternalServerError)
	}
}

func main() {
	router := mux.NewRouter()

	// Endpoints
	router.HandleFunc("/usuario", GetUsuario).Methods("GET")

	log.Printf("Iniciando programa na porta 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}

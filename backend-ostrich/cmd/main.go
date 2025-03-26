package main

import (
	"log"
	"net/http"

	"github.com/YKauan/ostrich-nostr-client/cmd/handlers"
)

// Middleware para permitir CORS
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")                            // Permite qualquer origem
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")          // Métodos permitidos
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization") // Cabeçalhos permitidos

		// Responder diretamente a requisições OPTIONS (preflight)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	// Configura as rotas
	mux.HandleFunc("/generate-keys", handlers.GenerateKeysHandler)
	mux.HandleFunc("/con", handlers.ConnectToRelayHandler)

	// Aplica o middleware CORS
	handlerWithCORS := enableCORS(mux)

	// Inicia o servidor na porta 8080
	log.Println("Servidor iniciado na porta 8080...")
	if err := http.ListenAndServe(":8080", handlerWithCORS); err != nil {
		log.Fatal(err)
	}
}

package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/YKauan/ostrich-nostr-client/cmd/service"
)

func ConnectToRelayHandler(w http.ResponseWriter, r *http.Request) {
	// Recupera a chave pública da URL (por exemplo, /feed?npub=chave_publica)
	npub := r.URL.Query().Get("npub")
	if npub == "" {
		http.Error(w, "Chave pública não fornecida", http.StatusBadRequest)
		return
	}

	// Chama a função que conecta ao relay e obtém o feed de publicações
	feed, err := service.ConnectToRelay(npub)
	if err != nil {
		http.Error(w, "Erro ao carregar feed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Retorna o feed como JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(feed)
}

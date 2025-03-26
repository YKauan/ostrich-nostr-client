package handlers

import (
    "encoding/json"
    "net/http"

    "github.com/YKauan/ostrich-nostr-client/cmd/service"
)

func GenerateKeysHandler(w http.ResponseWriter, r *http.Request) {
    privateKey, publicKey := service.GenerateKeys()
    response := map[string]string{
        "privateKey": privateKey,
        "publicKey":  publicKey,
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}
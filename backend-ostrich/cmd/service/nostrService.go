package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/nbd-wtf/go-nostr"
	"github.com/nbd-wtf/go-nostr/nip19"
)

var relays = []string{
	"wss://relay.damus.io",   // Exemplo de relay
	"wss://relay.nostr.band", // Outro exemplo de relay

}

type AuthorProfile struct {
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	About       string `json:"about"`
	Picture     string `json:"picture"`
	Banner      string `json:"banner"`
	Website     string `json:"website"`
	Nip05       string `json:"nip05"`
	Lud06       string `json:"lud06"`
	Lud16       string `json:"lud16"`
}

func GenerateKeys() (string, string) {
	// Gera a chave privada
	sk := nostr.GeneratePrivateKey()

	// Gera a chave pública a partir da chave privada
	pk, _ := nostr.GetPublicKey(sk)

	// Codifica as chaves no formato NIP-19
	nsec, _ := nip19.EncodePrivateKey(sk)
	npub, _ := nip19.EncodePublicKey(pk)

	return nsec, npub
}

// retorna os eventos do feed
func ConnectToRelay(npub string) ([]map[string]interface{}, error) {
	var feed []map[string]interface{}

	// Conexão com o relay
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	relay, err := nostr.RelayConnect(ctx, "wss://relay.damus.io") // Troque pelo seu relay se necessário
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao relay: %v", err)
	}

	// Decodifica a chave pública e prepara o filtro
	var filters nostr.Filters
	if _, _, err := nip19.Decode(npub); err == nil {
		// Filtro para pegar eventos de notas de texto (Kind 1)
		filters = []nostr.Filter{{
			Kinds: []int{nostr.KindTextNote}, // Tipo de evento (notas de texto)
			Limit: 50,                        // Limite de publicações
		}}
	} else {
		// Se não conseguir decodificar a chave pública, faz um filtro sem autor
		filters = []nostr.Filter{{
			Kinds: []int{nostr.KindTextNote}, // Tipo de evento
			Limit: 50,                        // Limite de publicações
		}}
	}

	// Assina os eventos com o filtro
	sub, err := relay.Subscribe(ctx, filters)
	if err != nil {
		return nil, fmt.Errorf("erro ao assinar os eventos: %v", err)
	}

	// Limite de eventos que vamos pegar antes de cancelar a inscrição
	maxEvents := 50
	eventCount := 0

	// Recebe os eventos e armazena no feed
	for ev := range sub.Events {
		if eventCount >= maxEvents {
			break
		}

		// Pega as informações do autor (nome e imagem) usando a chave pública
		name, image, err := GetAuthorInfoFromRelay(ev.PubKey)
		if err != nil {
			// Se não conseguir pegar o nome e imagem, coloca valores default
			name, image = "Autor Desconhecido", ""
		}

		// Cria um mapa com as informações do evento e do autor
		eventInfo := map[string]interface{}{
			"content":      ev.Content,
			"authorPubKey": ev.PubKey,    // A chave pública do autor
			"authorName":   name,         // Nome do autor
			"authorImage":  image,        // Imagem do autor
			"timestamp":    ev.CreatedAt, // O timestamp da publicação
			"tags":         ev.Tags,      // As tags associadas ao evento
		}

		// Adiciona o evento no feed
		feed = append(feed, eventInfo)
		eventCount++
	}

	// Unsubscribe do relay
	sub.Unsub()

	// Retorna o feed de publicações
	if len(feed) == 0 {
		return nil, fmt.Errorf("nenhum evento encontrado")
	}

	return feed, nil
}

func GetAuthorInfoFromRelay(npub string) (string, string, error) {
	// Conectar ao relay
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	relay, err := nostr.RelayConnect(ctx, "wss://relay.damus.io")
	if err != nil {
		return "", "", fmt.Errorf("erro ao conectar ao relay: %v", err)
	}

	// Definir o filtro para capturar eventos de tipo "Kind 0" (metadados do perfil)
	filters := nostr.Filters{
		{
			Kinds:   []int{0},       // Tipo de evento "Kind 0"
			Authors: []string{npub}, // Chave pública do autor
			Limit:   1,              // Limite para pegar apenas um evento
		},
	}

	// Inscrever-se para obter os eventos
	sub, err := relay.Subscribe(ctx, filters)
	if err != nil {
		return "", "", fmt.Errorf("erro ao assinar eventos: %v", err)
	}
	defer sub.Unsub()

	// Esperar pelo evento de perfil
	for ev := range sub.Events {
		if ev.Kind == 0 {
			// O evento contém informações de perfil
			// Por exemplo, o conteúdo do evento pode ser o nome do autor
			name := ev.Content // Assume-se que o conteúdo do evento seja o nome
			image := ""        // Você pode pegar a imagem dos tags, se disponível

			// Tenta desserializar a string JSON para um objeto de perfil
			var profile AuthorProfile
			if err := json.Unmarshal([]byte(ev.Content), &profile); err == nil {
				// Agora você pode acessar o nome, imagem e outras propriedades
				name = profile.DisplayName
				image = profile.Picture
			} else {
				// Se o JSON não puder ser desserializado, tenta usar o nome direto
				name = ev.Content
			}

			return name, image, nil
		}
	}

	return "", "", fmt.Errorf("não foi possível encontrar informações do autor")
}

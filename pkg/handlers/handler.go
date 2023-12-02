package handlers

import (
	"cloud.google.com/go/firestore"
)

type Handler struct {
	Client *firestore.Client
}

func New(client *firestore.Client) *Handler {
	return &Handler{
		Client: client,
	}

}

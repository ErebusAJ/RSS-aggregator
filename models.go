package main

import (
	"time"

	"github.com/ErebusAJ/rssagg/internal/database"
	"github.com/google/uuid"
)

type Users struct{
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAit time.Time `json:"updated_at"`
}

func databaseUserToUser(user database.User) Users{
	return Users{
		ID: user.ID,
		Name: user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAit: user.UpdatedAt,
	}
}
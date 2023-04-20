// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package db

import (
	"time"
)

type Chat struct {
	ID               string
	UserID           string
	InitialMessage   string
	Status           string
	TokenUsage       int16
	Model            string
	ModelMaxTokens   int16
	Temperature      string
	TopP             string
	N                int16
	Stop             string
	MaxTokens        int16
	PresencePenalty  string
	FrequencyPenalty string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type Message struct {
	ID        string
	ChatID    string
	Role      string
	Content   string
	Tokens    int16
	Model     string
	Erased    bool
	OrderMsg  int16
	CreatedAt time.Time
}

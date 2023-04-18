package entity

import (
	"errors"

	"github.com/google/uuid"
)

type ChatConfig struct {
	Model            *Model
	Temperature      float32
	TopP             float32
	N                int // number of mesaages to generate
	Stop             []string
	MaxTokens        int // max number of tokens to generate
	PresencePenalty  float32
	FrequencyPenalty float32
}

type Chat struct {
	ID                   string
	UserID               string
	InitialSystemMessage *Message
	Messages             []*Message
	ErasedMessages       []*Message
	Status               string
	TokenUsage           int
	Config               *ChatConfig
}

func NewChat(userID string, InitialSystemMessage *Message, config *ChatConfig) (*Chat, error) {
	chat := &Chat{
		ID:                   uuid.New().String(),
		UserID:               userID,
		InitialSystemMessage: InitialSystemMessage,
		Status:               "active",
		Config:               config,
		TokenUsage:           0,
	}
	chat.AddMessage(InitialSystemMessage)

	if err := chat.Validate(); err != nil {
		return nil, err
	}
	return chat, nil
}

func (c *Chat) AddMessage(m *Message) error {
	if c.Status == "ended" {
		return errors.New("chat is ended")
	}
	for {
		if c.Config.Model.getMaxTokens() >= m.getQuantityTokens()+c.TokenUsage {
			c.Messages = append(c.Messages, m)
			c.RefreshTokenUsage()
			break
		}
		c.ErasedMessages = append(c.ErasedMessages, c.Messages[0])
		c.Messages = c.Messages[1:]
		c.RefreshTokenUsage()
	}
	return nil
}

func (c *Chat) getMessages() []*Message {
	return c.Messages
}

func (c *Chat) countMessages() int {
	return len(c.Messages)
}

func (c *Chat) endChat() {
	c.Status = "ended"
}

func (c *Chat) isEnded() bool {
	return c.Status == "ended"
}

func (c *Chat) isStarted() bool {
	return c.Status == "active"
}

func (c *Chat) startChat() {
	c.Status = "active"
}

func (c *Chat) Validate() error {
	if c.UserID == "" {
		return errors.New("user id is empty")
	}
	if c.Status != "started" && c.Status != "ended" {
		return errors.New("invalid status")
	}
	if c.Config.Temperature < 0.0 || c.Config.Temperature > 5.0 {
		return errors.New("invalid temperature")
	}
	return nil
}

func (c *Chat) RefreshTokenUsage() {
	c.TokenUsage = 0
	for m := range c.Messages {
		c.TokenUsage += c.Messages[m].getQuantityTokens()
	}
}

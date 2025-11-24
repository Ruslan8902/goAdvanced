package auth

import (
	"math/rand"

	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	UserID           uint   `json:"userId"`
	SessionID        string `json:"sessionId"`
	Token            string `json:"token"`
	ConfirmationCode string `json:"confirmationCode"`
}

func NewSession(userId uint) *Session {
	session := &Session{
		UserID: userId,
	}
	session.GenerateConfirmationCode()
	session.GenerateSessionID()
	return session
}

func (session *Session) GenerateSessionID() {
	session.SessionID = RandStringRunes(16, digitRunes)
}

func (session *Session) GenerateConfirmationCode() {
	session.SessionID = RandStringRunes(16, digitRunes)
}

var symbolsRunes = []rune("abcdefghijklmnoprstuvwxyzABCDEFGHIJKLMNOPRSTUVWXYZ0123456789")
var digitRunes = []rune("0123456789")

func RandStringRunes(n int, symbols []rune) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = symbols[rand.Intn(len(symbols))]
	}
	return string(b)
}

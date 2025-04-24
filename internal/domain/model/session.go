package model

type SessionSecret string

type Session struct {
	secret SessionSecret
}

func NewSession(secret string) *Session {
	return &Session{
		secret: SessionSecret(secret),
	}
}

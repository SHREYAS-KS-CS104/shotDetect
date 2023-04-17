package models

import (
	"database/sql"
	"fmt"

	"github.com/SHREYAS-KS-CS104/shotDetect/rand"
)

type Session struct {
	ID     int
	UserID int
	// Token is only set when creating a new session.
	// When looking up a seesion this will be left empty
	// as we only store the hash a session token
	// in our database and we cannot reverse it into a raw token.
	Token     string
	TokenHash string
}

type SessionService struct {
	DB *sql.DB
}

func (ss *SessionService) Create(userID int) (*Session, error) {
	token, err := rand.SessionToken()
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}
	//TODO: Hash the session token
	session := Session{
		UserID: userID,
		Token:  token,
		// TODO: Set the toeknHash
	}
	// TODO: Store the session in our DB
	return &session, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	// TODO Implement SesionService.User
	return nil, nil
}

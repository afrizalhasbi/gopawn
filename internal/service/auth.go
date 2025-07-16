package service

import (
	"database/sql"
	"fmt"
	"gopawn/internal/data/payload"
	"log"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	DB *sql.DB
}

func (s *AuthService) Register(reg *payload.Login) error {
	var uuid = uuid.New().String()
	var now = time.Now().Format(time.RFC3339)
	var pwdHash, err = bcrypt.GenerateFromPassword([]byte(reg.Password), 12)
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		return err
	} else {
		tx, err := s.DB.Begin()
		if err != nil {
			return fmt.Errorf("failed to begin transaction: %w", err)
		}

		defer tx.Rollback()
		_, err = tx.Exec(
			"INSERT INTO users (uuid, name, created, updated, elo, games) VALUES ($1, $2, $3, $4, $5, $6)",
			uuid, reg.Name, now, now, 1500, 0,
		)
		if err != nil {
			return fmt.Errorf("failed to create user: %w", err)
		}

		_, err = tx.Exec(
			"INSERT INTO auth (name, email, password) VALUES ($1, $2, $3)",
			reg.Name, reg.Email, string(pwdHash),
		)
		if err != nil {
			return fmt.Errorf("failed to create user: %w", err)
		}

		tx.Commit()
		return nil
	}
}

func (s *AuthService) Login(login *payload.Login) error {
	var hashedPassword string
	err := s.DB.QueryRow("SELECT password FROM auth WHERE email = $1", login.Email).Scan(&hashedPassword)
	if err != nil {
		return err
	}
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(login.Password))
}

func (s *AuthService) Logout() error {
	return nil
}

func (s *AuthService) ForgotPassword(reset *payload.Reset) error {
	return nil
}

func (s *AuthService) ResetPassword(reset *payload.Reset) error {
	return nil
}

func (s *AuthService) Delete() {}

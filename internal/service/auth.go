package service

import (
	"database/sql"
	"fmt"
	"gopawn/internal/data/payload"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	DB        *sql.DB
	SecretKey []byte
}

type Claims struct {
	Email string
	jwt.RegisteredClaims
}

func (s *AuthService) Register(reg *payload.Register) error {
	var uuid = uuid.New().String()
	var now = time.Now().Format(time.RFC3339)
	var pwdHash, err = bcrypt.GenerateFromPassword([]byte(reg.Password), 12)
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		return err
	} else {
		tx, err := s.DB.Begin()
		if err != nil {
			return fmt.Errorf("Failed to register user")
		}

		defer tx.Rollback()
		_, err = tx.Exec(
			"INSERT INTO users (uuid, name, created, updated, elo, games) VALUES ($1, $2, $3, $4, $5, $6)",
			uuid, reg.Name, now, now, 1500, 0,
		)
		if err != nil {
			return fmt.Errorf("Failed to register user.")
		}

		_, err = tx.Exec(
			"INSERT INTO auth (name, email, password) VALUES ($1, $2, $3)",
			reg.Name, reg.Email, string(pwdHash),
		)
		if err != nil {
			return fmt.Errorf("Failed to register user.")
		}
		tx.Commit()
		return nil
	}
}

func (s *AuthService) Login(login *payload.Login) (string, error) {
	err := s.__ValidateCredentials(login.Email, login.Name)
	if err != nil {
		return "", fmt.Errorf("Invalid credentials")
	} else {
		jwtTokenString, err := s.__GenerateJwt(login.Email)
		if err != nil {
			return "", fmt.Errorf("Failed to generate authentication token")
		} else {
			return jwtTokenString, nil
		}
	}
}

func (s *AuthService) Logout() error {
	return nil
}

func (s *AuthService) ForgotPassword(reset *payload.ForgotPassword) error {
	err := s.__ValidateCredentials(reset.Email, reset.Name)
	if err != nil {
		return err
	} else {
		s.__SendEmail(reset.Email)
		return nil
	}
}

func (s *AuthService) ResetPassword(reset *payload.ResetPassword) error {
	if reset.NewPassword == reset.NewPasswordCopy {
		tx, err := s.DB.Begin()
		if err != nil {
			return fmt.Errorf("Failed to reset password")
		}
		defer tx.Rollback()
		_, err = tx.Exec(
			"INSERT INTO auth (password) VALUES ($1) WHERE `email` = ($2)",
			reset.NewPassword, reset.Email,
		)
		if err != nil {
			return fmt.Errorf("Failed to reset password")
		}
		tx.Commit()
	}
	return nil
}

func (s *AuthService) Delete(delete *payload.Delete) error {
	err := s.__ValidateCredentials(delete.Email, delete.Name)
	if err != nil {
		return fmt.Errorf("Invalid credentials")
	} else {
		tx, err := s.DB.Begin()
		if err != nil {
			return fmt.Errorf("Failed to delete user")
		}
		defer tx.Rollback()
		_, err = tx.Exec("DROP FROM auth WHERE email = ($1)", delete.Email)
		if err != nil {
			return fmt.Errorf("Failed to delete user")
		}
		tx.Commit()
	}
	return nil
}

func (s *AuthService) __GenerateJwt(email string) (string, error) {
	claims := Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Admin",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtTokenString, err := jwtToken.SignedString(s.SecretKey)
	if err != nil {
		return "", err
	} else {
		return jwtTokenString, nil
	}
}

func (s *AuthService) __SendEmail(email string) {
	// not implemented for realsies
}

func (s *AuthService) __ValidateCredentials(email string, password string) error {
	var hashedPassword string
	err := s.DB.QueryRow("SELECT password FROM auth WHERE email = $1", email).Scan(&hashedPassword)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err
}

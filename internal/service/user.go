package service

import (
	"database/sql"
	"gopawn/internal/data/schema"
)

type UserService struct {
	DB *sql.DB
}

func (s *UserService) UpdateGame(user schema.User) {
	s.DB.Query("UPDATE users () VALUES ($2 $3 $4 $5 $6 )", user.Name, user.Updated)
}

func (s *UserService) Delete(user schema.User) {}

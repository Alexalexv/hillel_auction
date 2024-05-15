package repository

import (
	"database/sql"
	"errors"
	"hillel_auction/db"
	"hillel_auction/logger"
	"hillel_auction/repository/models"
)

type UsersRepository interface {
	Insert(user *models.User) (*models.User, error)
	GetByEmailPassword(email string, pass string) (*models.User, error)
}

type UsersRepos struct {
	db  *db.Database
	log *logger.Logger
}

func NewUsersRepository(db *db.Database, log *logger.Logger) UsersRepository {
	return &UsersRepos{
		db:  db,
		log: log,
	}
}

func (ur UsersRepos) Insert(user *models.User) (*models.User, error) {
	var id int

	err := ur.db.QueryRowx("INSERT INTO users (email, password) VALUES($1, $2) RETURNING id",
		user.Email, user.Password).Scan(&id)
	if err != nil {
		ur.log.Error("failed to write in db", err)
		return nil, err
	}

	user.ID = id

	return user, nil
}

func (ur UsersRepos) GetByEmailPassword(email string, pass string) (*models.User, error) {
	u := &models.User{}

	err := ur.db.Get(u, "SELECT * FROM users WHERE email=$1 AND password=$2", email, pass)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		ur.log.Error("failed to get user", err)
		return nil, err
	}
	return u, nil
}

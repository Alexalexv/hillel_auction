package services

import (
	"hillel_auction/logger"
	"hillel_auction/repository"
	"hillel_auction/repository/models"
)

type UserService struct {
	userRepos repository.UsersRepository
	log       *logger.Logger
}

func NewUserService(log *logger.Logger, userRepos repository.UsersRepository) *UserService {
	return &UserService{
		userRepos: userRepos,
		log:       log,
	}
}

func (us *UserService) CreateUser(user *models.User) (*models.User, error) {
	return us.userRepos.Insert(user)
}

func (us *UserService) FindUserByEmailPass(email string, pass string) (*models.User, error) {
	return us.userRepos.GetByEmailPassword(email, pass)
}

package service

import "pipelines/internal/model"

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetUsers() []model.User {

	return []model.User{
		{
			ID:   1,
			Name: "CI",
		},
		{
			ID:   2,
			Name: "CD",
		},
	}
}

package service

import (
	"api/database/titlerepo"
	"api/model"
	"api/user"
	"context"
	"go.uber.org/zap"
)

type UserRepo interface {
	InsertData(ctx context.Context, userModel model.User, locationModel model.Location, loginModel model.Login, pictureModel model.Picture) error
	SelectData(ctx context.Context) (string, error)
}

type UserService struct {
	logger *zap.SugaredLogger
	cli    User.Client
	repo   titlerepo.Repository
}

func NewUserService(logger *zap.SugaredLogger, cli User.Client, repo titlerepo.Repository) UserService {
	return UserService{
		logger: logger,
		cli:    cli,
		repo:   repo,
	}
}

func (s UserService) InsertInformation(ctx context.Context) error {
	userData, locationData, loginData, pictureData, err := s.cli.GetInformation(ctx)
	if err != nil {
		return err
	}

	if err = s.repo.InsertData(ctx, userData, locationData, loginData, pictureData); err != nil {
		return err
	}

	return nil
}

func (s UserService) SelectInformation(ctx context.Context) (string, error) {
	data, err := s.repo.SelectData(ctx)
	if err != nil {
		return "", err
	}

	return data, nil
}

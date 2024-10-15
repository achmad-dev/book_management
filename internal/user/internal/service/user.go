package service

import (
	"context"
	"errors"

	"github.com/achmad-dev/internal/user/api/dto"
	"github.com/achmad-dev/internal/user/internal/domain"
	"github.com/achmad-dev/internal/user/internal/repository"
	"github.com/achmad-dev/internal/user/internal/util"
	"github.com/sirupsen/logrus"
)

type UserService interface {
	GetUser(ctx context.Context, id string) (*domain.User, error)
	CreateUser(ctx context.Context, signUpRequest dto.SignUpRequest) error
	SignIn(ctx context.Context, signInRequest dto.SignInRequest) (string, error)
	GetUserByUsername(ctx context.Context, username string) (*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) error
	DeleteUser(ctx context.Context, id string) error
}

type userServiceImpl struct {
	userRepo   repository.UserRepository
	bcryptUtil util.BcryptUtil
	log        *logrus.Logger
	secret     string
}

// GetUserByUsername implements UserService.
func (u *userServiceImpl) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	return u.userRepo.GetUserByUsername(ctx, username)
}

// SignIn implements UserService.
func (u *userServiceImpl) SignIn(ctx context.Context, signInRequest dto.SignInRequest) (string, error) {
	user, err := u.userRepo.GetUserByUsername(ctx, signInRequest.Username)
	if err != nil {
		return "", err
	}
	isUser := u.bcryptUtil.CheckPasswordHash(signInRequest.Password, user.Password)
	if !isUser {
		return "", errors.New("invalid password")
	}
	token, err := util.GenerateToken(user.Username, u.secret)
	if err != nil {
		u.log.Errorf("failed to generate token: %v", err)
		return "", err
	}
	return token, nil
}

// CreateUser implements userService.
func (u *userServiceImpl) CreateUser(ctx context.Context, signUpRequest dto.SignUpRequest) error {
	hashedPassword, err := u.bcryptUtil.HashPassword(signUpRequest.Password)
	if err != nil {
		u.log.Errorf("failed to hash password: %v", err)
		return err
	}
	user := &domain.User{
		Username: signUpRequest.Username,
		Password: hashedPassword,
		Role:     signUpRequest.Role,
	}
	err = u.userRepo.CreateUser(ctx, user)
	if err != nil {
		u.log.Errorf("failed to create user: %v", err)
		return err
	}
	return nil
}

// DeleteUser implements userService.
func (u *userServiceImpl) DeleteUser(ctx context.Context, id string) error {
	panic("unimplemented")
}

// GetUser implements userService.
func (u *userServiceImpl) GetUser(ctx context.Context, id string) (*domain.User, error) {
	panic("unimplemented")
}

// UpdateUser implements userService.
func (u *userServiceImpl) UpdateUser(ctx context.Context, user *domain.User) error {
	panic("unimplemented")
}

func NewUserService(userRepo repository.UserRepository, bcryptUtil util.BcryptUtil, log *logrus.Logger, secret string) UserService {
	return &userServiceImpl{userRepo: userRepo, bcryptUtil: bcryptUtil, log: log, secret: secret}
}

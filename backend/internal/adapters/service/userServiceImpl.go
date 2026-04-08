package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/undndnwnkk/go-react-todoapp/internal/adapters/service/helpers"
	"github.com/undndnwnkk/go-react-todoapp/internal/core/domain"
	"github.com/undndnwnkk/go-react-todoapp/internal/core/ports"
)

type UserServiceImpl struct {
	repo ports.UserRepository
}

func (u *UserServiceImpl) GetAll(ctx context.Context) ([]domain.User, error) {
	users, err := u.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserServiceImpl) Register(ctx context.Context, request domain.UserCreateRequest) (domain.UserIdResponse, error) {
	if err := helpers.CheckUserCreateRequest(request); err != nil {
		return domain.UserIdResponse{}, err
	}

	hashedPassword, err := helpers.HashPassword(request.Password)
	if err != nil {
		return domain.UserIdResponse{}, err
	}

	request.Password = hashedPassword

	user, err := u.repo.Create(ctx, request)
	if err != nil {
		return domain.UserIdResponse{}, err
	}

	return user, nil
}

func (u *UserServiceImpl) Login(ctx context.Context, request domain.UserLoginRequest) (domain.User, error) {
	user, err := u.repo.GetByEmail(ctx, request.Email)
	if err != nil {
		return domain.User{}, err
	}

	if !helpers.CheckPasswordHash(request.Password, user.PasswordHash) {
		return domain.User{}, domain.ErrInvalidPassword
	}

	return user, nil
}

func (u *UserServiceImpl) GetByID(ctx context.Context, id uuid.UUID) (domain.User, error) {
	user, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (u *UserServiceImpl) UpdateByID(ctx context.Context, id uuid.UUID, request domain.UserUpdateRequest) (domain.User, error) {
	user, err := u.repo.UpdateByID(ctx, id, request)
	if err != nil {
		return domain.User{}, err
	}

	return user, err
}

func (u *UserServiceImpl) DeleteByID(ctx context.Context, id uuid.UUID) error {
	if err := u.repo.DeleteByID(ctx, id); err != nil {
		return err
	}

	return nil
}

func NewUserService(r ports.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{repo: r}
}

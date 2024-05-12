package user_usecase

import (
	"auction_golang_concurrency/internal/entity/user_entity"
	"auction_golang_concurrency/internal/internal_error"
	"context"
)

type UserUseCase struct {
	UserRepository user_entity.UserRepositoryInterface
}

func NewUserUseCase(userRepository user_entity.UserRepositoryInterface) UserUseCaseInterface {
	return UserUseCase{
		UserRepository: userRepository,
	}
}

type UserOutputDTO struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type UserUseCaseInterface interface {
	FindByUserById(ctx context.Context, id string) (*UserOutputDTO, *internal_error.InternalError)
}

func (usecase UserUseCase) FindByUserById(ctx context.Context, id string) (*UserOutputDTO, *internal_error.InternalError) {

	userEntity, err := usecase.UserRepository.FindUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &UserOutputDTO{
		Id:   userEntity.Id,
		Name: userEntity.Name,
	}, nil
}

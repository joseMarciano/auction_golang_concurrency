package user_controller

import (
	"auction_golang_concurrency/configuration/rest_err"
	"auction_golang_concurrency/internal/usecase/user_usecase"
	"context"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUseCase user_usecase.UserUseCaseInterface
}

func NewUserController(userUseCase user_usecase.UserUseCaseInterface) *UserController {
	return &UserController{
		UserUseCase: userUseCase,
	}
}

func (uc *UserController) FindUserById(c *gin.Context) {
	userId := c.Param("id")
	output, err := uc.UserUseCase.FindByUserById(context.TODO(), userId)
	if err != nil {
		errRest := rest_err.NewNotFoundError("user not found")
		c.JSON(errRest.Status, errRest)
		return
	}

	c.JSON(200, output)
}

package auction_controller

import (
	"auction_golang_concurrency/configuration/rest_err"
	"auction_golang_concurrency/internal/usecase/auction_usecase"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuctionController struct {
	AuctionUseCase *auction_usecase.AuctionUseCase
}

func NewAuctionController(userUseCase *auction_usecase.AuctionUseCase) *AuctionController {
	return &AuctionController{
		AuctionUseCase: userUseCase,
	}
}

func (useCase *AuctionController) CreateAuction(c *gin.Context) {
	var auctionInputDto auction_usecase.AuctionInputDto

	if err := c.ShouldBindJSON(&auctionInputDto); err != nil {
		restErr := rest_err.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	err := useCase.AuctionUseCase.CreateAuction(context.Background(), &auctionInputDto)
	if err != nil {
		restErr := rest_err.NewBadRequestError("error on create auction")
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusCreated, nil)
}

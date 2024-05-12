package bid_controller

import (
	"auction_golang_concurrency/configuration/rest_err"
	"auction_golang_concurrency/internal/usecase/bid_usecase"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BidController struct {
	BidUseCase bid_usecase.BidUseCaseInterface
}

func NewBidController(userUseCase bid_usecase.BidUseCaseInterface) *BidController {
	return &BidController{
		BidUseCase: userUseCase,
	}
}

func (useCase *BidController) CreateBid(c *gin.Context) {
	var auctionInputDto bid_usecase.BidInputDto

	if err := c.ShouldBindJSON(&auctionInputDto); err != nil {
		restErr := rest_err.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	output, err := useCase.BidUseCase.CreateBid(context.Background(), &auctionInputDto)
	if err != nil {
		restErr := rest_err.NewBadRequestError("error on create auction")
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusCreated, output)
}

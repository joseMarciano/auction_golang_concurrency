package bid_controller

import (
	"auction_golang_concurrency/configuration/rest_err"
	"context"
	"github.com/gin-gonic/gin"
)

func (useCase *BidController) FindBidByAuctionId(c *gin.Context) {
	auctionId := c.Param("id")
	output, err := useCase.BidUseCase.FindByAuctionId(context.TODO(), auctionId)
	if err != nil {
		errRest := rest_err.NewNotFoundError("auction not found")
		c.JSON(errRest.Status, errRest)
		return
	}

	c.JSON(200, output)
}

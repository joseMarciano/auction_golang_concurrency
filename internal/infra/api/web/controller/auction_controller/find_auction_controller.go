package auction_controller

import (
	"auction_golang_concurrency/configuration/rest_err"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (useCase *AuctionController) FindAuctionById(c *gin.Context) {
	auctionId := c.Param("id")
	output, err := useCase.AuctionUseCase.FindAuctionByID(context.TODO(), auctionId)
	if err != nil {
		errRest := rest_err.NewNotFoundError("auction not found")
		c.JSON(errRest.Status, errRest)
		return
	}

	c.JSON(200, output)
}

func (useCase *AuctionController) FindAuctions(c *gin.Context) {
	status, _ := strconv.Atoi(c.Query("status"))
	category := c.Query("category")
	productName := c.Query("productName")

	output, err := useCase.AuctionUseCase.FindAuctions(context.TODO(), status, category, productName)
	if err != nil {
		errRest := rest_err.NewBadRequestError("error on find auctions")
		c.JSON(errRest.Status, errRest)
		return
	}

	c.JSON(http.StatusOK, output)
}

func (useCase *AuctionController) FindWinningBidByAuctionId(c *gin.Context) {
	auctionId := c.Param("id")
	output, err := useCase.AuctionUseCase.FindAuctionByID(context.TODO(), auctionId)
	if err != nil {
		errRest := rest_err.NewNotFoundError("auction not found")
		c.JSON(errRest.Status, errRest)
		return
	}

	bidOutput, err := useCase.AuctionUseCase.FindWinningBidByAuctionId(context.TODO(), output.Id)
	if err != nil {
		errRest := rest_err.NewBadRequestError("error on find winning bid")
		c.JSON(errRest.Status, errRest)
		return
	}

	c.JSON(http.StatusOK, bidOutput)
}

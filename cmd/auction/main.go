package main

import (
	"auction_golang_concurrency/configuration/database/mongodb"
	"auction_golang_concurrency/internal/infra/api/web/controller/auction_controller"
	"auction_golang_concurrency/internal/infra/api/web/controller/bid_controller"
	"auction_golang_concurrency/internal/infra/api/web/controller/user_controller"
	"auction_golang_concurrency/internal/infra/database/auction"
	"auction_golang_concurrency/internal/infra/database/bid"
	"auction_golang_concurrency/internal/infra/database/user"
	"auction_golang_concurrency/internal/usecase/auction_usecase"
	"auction_golang_concurrency/internal/usecase/bid_usecase"
	"auction_golang_concurrency/internal/usecase/user_usecase"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
		return
	}

	database, err := mongodb.NewMongoDbConnection(context.TODO())
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		return
	}

	userController, bidController, auctionController := initDependencies(database)

	router := gin.Default()
	router.GET("/auctions", auctionController.FindAuctions)
	router.POST("/auctions", auctionController.CreateAuction)
	router.GET("/auction/winner/:id", auctionController.FindWinningBidByAuctionId)
	router.POST("/bid", bidController.CreateBid)
	router.GET("/bid/:id", bidController.FindBidByAuctionId)
	router.GET("/user/:id", userController.FindUserById)

	router.Run(":8080")
}

func initDependencies(database *mongo.Database) (
	userController *user_controller.UserController,
	bidController *bid_controller.BidController,
	auctionController *auction_controller.AuctionController) {

	userRepo := user.NewUserRepository(database)
	userUseCase := user_usecase.NewUserUseCase(userRepo)
	user_controller.NewUserController(userUseCase)

	auctionRepo := auction.NewAuctionRepository(database)
	bidRepo := bid.NewBidRepository(database, auctionRepo)
	auctionUseCase := auction_usecase.NewAuctionUseCase(auctionRepo, bidRepo)
	bidUseCase := bid_usecase.NewBidUseCase(bidRepo)

	return user_controller.NewUserController(userUseCase), bid_controller.NewBidController(bidUseCase), auction_controller.NewAuctionController(auctionUseCase)
	//userUseCase := user_usecase.NewUserUseCase(userRepo)
	//user_controller.NewUserController(userUseCase)
}

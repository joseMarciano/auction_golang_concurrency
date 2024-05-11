package auction

import (
	"auction_golang_concurrency/configuration/logger"
	"auction_golang_concurrency/internal/entity/auction_entity"
	"auction_golang_concurrency/internal/internal_error"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuctionEntityMongo struct {
	Id          string                          `bson:"_id"`
	ProductName string                          `bson:"productName"`
	Category    string                          `bson:"category"`
	Description string                          `bson:"description"`
	Condition   auction_entity.AuctionCondition `bson:"condition"`
	Status      auction_entity.AuctionStatus    `bson:"status"`
	CreatedAt   int64                           `bson:"createdAt"`
}

type AuctionRepository struct {
	Collection *mongo.Collection
}

func NewAuctionRepository(database *mongo.Database) *AuctionRepository {
	return &AuctionRepository{Collection: database.Collection("auctions")}
}

func (ar *AuctionRepository) CreateAuction(ctx context.Context, auction *auction_entity.Auction) *internal_error.InternalError {
	auctionMongo := &AuctionEntityMongo{
		Id:          auction.Id,
		ProductName: auction.ProductName,
		Category:    auction.Category,
		Description: auction.Description,
		Condition:   auction.Condition,
		Status:      auction.Status,
		CreatedAt:   auction.CreatedAt.Unix(),
	}

	_, err := ar.Collection.InsertOne(ctx, auctionMongo)
	if err != nil {
		logger.Error("error on create auction", err)
		return internal_error.NewInternalError(err.Error())
	}

	return nil

}

package auction

import (
	"auction_golang_concurrency/configuration/logger"
	"auction_golang_concurrency/internal/entity/auction_entity"
	"auction_golang_concurrency/internal/internal_error"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (ar *AuctionRepository) FindAuctionByID(ctx context.Context, id string) (*auction_entity.Auction, *internal_error.InternalError) {
	filter := bson.M{"_id": id}

	var auctionEntityMongo AuctionEntityMongo

	if err := ar.Collection.FindOne(ctx, filter).Decode(&auctionEntityMongo); err != nil {
		errMessage := fmt.Sprintf("error trying to findAuction with bid %s error %s", id, err.Error())
		logger.Error(errMessage, err)
		return nil, internal_error.NewInternalError(errMessage)
	}

	return &auction_entity.Auction{
		Id:          auctionEntityMongo.Id,
		ProductName: auctionEntityMongo.ProductName,
		Category:    auctionEntityMongo.Category,
		Description: auctionEntityMongo.Description,
		Condition:   auctionEntityMongo.Condition,
		Status:      auctionEntityMongo.Status,
		CreatedAt:   time.Unix(auctionEntityMongo.CreatedAt, 0),
	}, nil
}

func (ar *AuctionRepository) FindAuctions(
	ctx context.Context,
	status auction_entity.AuctionStatus,
	category,
	productName string,
) ([]auction_entity.Auction, *internal_error.InternalError) {

	filter := bson.M{}

	if status != 0 {
		filter["status"] = status
	}

	if category != "" {
		filter["category"] = category
	}

	if productName != "" {
		filter["productName"] = primitive.Regex{
			Pattern: productName,
			Options: "i",
		}
	}

	cursor, err := ar.Collection.Find(ctx, filter)
	if err != nil {
		errMessage := fmt.Sprintf("error trying to findAuctions with status %s category %s productName %s error %s",
			status, category, productName, err.Error())
		logger.Error(errMessage, err)
		return nil, internal_error.NewInternalError(errMessage)
	}

	defer cursor.Close(ctx)

	var auctions []AuctionEntityMongo
	err = cursor.All(ctx, &auctions)
	if err != nil {
		errMessage := fmt.Sprintf("error trying to findAuctions with status %s category %s productName %s error %s",
			status, category, productName, err.Error())
		logger.Error(errMessage, err)
		return nil, internal_error.NewInternalError(errMessage)
	}

	var auctionEntities []auction_entity.Auction
	for _, auction := range auctions {
		auctionEntities = append(auctionEntities, auction_entity.Auction{
			Id:          auction.Id,
			ProductName: auction.ProductName,
			Category:    auction.Category,
			Description: auction.Description,
			Condition:   auction.Condition,
			Status:      auction.Status,
			CreatedAt:   time.Unix(auction.CreatedAt, 0),
		})
	}

	return auctionEntities, nil
}

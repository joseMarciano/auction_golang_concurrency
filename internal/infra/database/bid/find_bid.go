package bid

import (
	"auction_golang_concurrency/configuration/logger"
	"auction_golang_concurrency/internal/entity/bid_entity"
	"auction_golang_concurrency/internal/internal_error"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func (bd *BidRepository) FindByAuctionId(ctx context.Context, auctionId string) ([]bid_entity.Bid, *internal_error.InternalError) {
	filter := bson.M{"auction_id": auctionId}

	cursor, err := bd.Collection.Find(ctx, filter)
	if err != nil {
		errorMessage := fmt.Sprintf("error trying to find bids by auctionId %s", auctionId)
		logger.Error(errorMessage, err)
		return nil, internal_error.NewInternalError(errorMessage)
	}

	var bidsEntityMongo []BidEntityMongo
	if err = cursor.All(ctx, &bidsEntityMongo); err != nil {
		errorMessage := fmt.Sprintf("error trying to find bids by auctionId %s", auctionId)
		logger.Error(errorMessage, err)
		return nil, internal_error.NewInternalError(errorMessage)
	}

	var bidsEntity []bid_entity.Bid
	for _, bidEntityMongo := range bidsEntityMongo {
		bidsEntity = append(bidsEntity, bid_entity.Bid{
			Id:        bidEntityMongo.Id,
			UserId:    bidEntityMongo.UserId,
			AuctionId: bidEntityMongo.AuctionId,
			Amount:    bidEntityMongo.Amount,
			CreatedAt: time.Unix(bidEntityMongo.CreatedAt, 0),
		})
	}

	return bidsEntity, nil
}

func (bd *BidRepository) FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*bid_entity.Bid, *internal_error.InternalError) {
	filter := bson.M{"auction_id": auctionId}

	var bidEntityMongo BidEntityMongo
	opts := options.FindOne().SetSort(bson.D{{"amount", -1}}) // sort by biggest amount

	if err := bd.Collection.FindOne(ctx, filter, opts).Decode(&bidEntityMongo); err != nil {
		errorMessage := fmt.Sprintf("error trying to find winning bid by auctionId %s", auctionId)
		logger.Error(errorMessage, err)
		return nil, internal_error.NewInternalError(errorMessage)
	}

	return &bid_entity.Bid{
		Id:        bidEntityMongo.Id,
		UserId:    bidEntityMongo.UserId,
		AuctionId: bidEntityMongo.AuctionId,
		Amount:    bidEntityMongo.Amount,
		CreatedAt: time.Unix(bidEntityMongo.CreatedAt, 0),
	}, nil
}

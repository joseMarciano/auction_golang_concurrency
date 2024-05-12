package bid

import (
	"auction_golang_concurrency/configuration/logger"
	"auction_golang_concurrency/internal/entity/auction_entity"
	"auction_golang_concurrency/internal/entity/bid_entity"
	"auction_golang_concurrency/internal/infra/database/auction"
	"auction_golang_concurrency/internal/internal_error"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

type BidEntityMongo struct {
	Id        string  `bson:"_id"`
	UserId    string  `bson:"userId"`
	AuctionId string  `bson:"auctionId"`
	Amount    float64 `bson:"amount"`
	CreatedAt int64   `bson:"createdAt"`
}

type BidRepository struct {
	Collection        *mongo.Collection
	AuctionRepository *auction.AuctionRepository
}

func NewBidRepository(database *mongo.Database, auctionRepository *auction.AuctionRepository) *BidRepository {
	return &BidRepository{Collection: database.Collection("bids"), AuctionRepository: auctionRepository}
}

func (bd *BidRepository) CreateBid(ctx context.Context, bidEntities []bid_entity.Bid) *internal_error.InternalError {
	var wg sync.WaitGroup

	for _, bid := range bidEntities {
		wg.Add(1)

		go func(bidValue bid_entity.Bid) {
			defer wg.Done()
			auctionEntity, err := bd.AuctionRepository.FindAuctionByID(ctx, bidValue.AuctionId)
			if err != nil {
				logger.Error("error trying to find auctionEntity by id ", err)
				return
			}

			if auctionEntity.Status != auction_entity.Active {
				return
			}

			var bidEntity = BidEntityMongo{
				Id:        bidValue.Id,
				UserId:    bidValue.UserId,
				AuctionId: bidValue.AuctionId,
				Amount:    bidValue.Amount,
				CreatedAt: bidValue.CreatedAt.Unix(),
			}

			if _, err := bd.Collection.InsertOne(ctx, bidEntity); err != nil {
				logger.Error("error trying to insert bidEntity ", err)
				return
			}

		}(bid)
	}

	wg.Wait()
	return nil
}

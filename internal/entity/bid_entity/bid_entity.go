package bid_entity

import (
	"auction_golang_concurrency/internal/internal_error"
	"context"
	"github.com/google/uuid"
	"time"
)

type Bid struct {
	Id        string
	UserId    string
	AuctionId string
	Amount    float64
	CreatedAt time.Time
}

func NewBid(userId string, auctionId string, amount float64) *Bid {
	return &Bid{
		Id:        uuid.New().String(),
		UserId:    userId,
		AuctionId: auctionId,
		Amount:    amount,
		CreatedAt: time.Now(),
	}
}

type BidRepositoryInterface interface {
	CreateBid(ctx context.Context, bidEntities []Bid) *internal_error.InternalError
	FindByAuctionId(ctx context.Context, auctionId string) ([]Bid, *internal_error.InternalError)
	FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*Bid, *internal_error.InternalError)
}

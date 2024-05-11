package bid_usecase

import (
	"auction_golang_concurrency/internal/entity/bid_entity"
	"auction_golang_concurrency/internal/internal_error"
	"context"
	"time"
)

type BidInputDto struct {
	UserId    string  `json:"userId"`
	AuctionId string  `json:"auctionId"`
	Amount    float64 `json:"amount"`
}

type BidOutputDto struct {
	Id        string    `json:"id"`
	UserId    string    `json:"userId"`
	AuctionId string    `json:"auctionId"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"createdAt" time_format:"2006-01-02 15:04:05"`
}

type BidUseCase struct {
	BidRepository bid_entity.BidRepositoryInterface
}

func (useCase *BidUseCase) CreateBid(ctx context.Context, input *BidInputDto) (*BidOutputDto, *internal_error.InternalError) {

	return nil, nil
}

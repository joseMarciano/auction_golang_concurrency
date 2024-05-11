package auction_usecase

import (
	"auction_golang_concurrency/internal/entity/auction_entity"
	"auction_golang_concurrency/internal/entity/bid_entity"
	"auction_golang_concurrency/internal/internal_error"
	"auction_golang_concurrency/internal/usecase/bid_usecase"
	"context"
	"time"
)

type AuctionInputDto struct {
	ProductName string                          `json:"productName"`
	Category    string                          `json:"category"`
	Description string                          `json:"description"`
	Condition   auction_entity.AuctionCondition `json:"condition"`
}

type AuctionOutputDto struct {
	Id          string                          `json:"id"`
	ProductName string                          `json:"productName"`
	Category    string                          `json:"category"`
	Description string                          `json:"description"`
	Condition   auction_entity.AuctionCondition `json:"condition"`
	Status      auction_entity.AuctionStatus    `json:"status"`
	CreatedAt   time.Time                       `json:"createdAt" time_format:"2006-01-02 15:04:05"`
}

type WinningInfoOutputDto struct {
	Auction *AuctionOutputDto         `json:"auction"`
	Bid     *bid_usecase.BidOutputDto `json:"bid"`
}

type AuctionUseCase struct {
	AuctionRepository auction_entity.AuctionRepositoryInterface
	BidRepository     bid_entity.BidRepositoryInterface
}

//type AuctionUseCase interface {
//	CreateAuction(ctx context.Context, input *AuctionInputDto) (*AuctionOutputDto, *internal_error.InternalError)
//	FindAuctionByID(ctx context.Context, id string) (*AuctionOutputDto, *internal_error.InternalError)
//	FindAuctions(ctx context.Context, status int, category, productName string) ([]AuctionOutputDto, *internal_error.InternalError)
//}

func (useCase *AuctionUseCase) CreateAuction(ctx context.Context, input *AuctionInputDto) *internal_error.InternalError {
	auction := auction_entity.CreateAuction(input.ProductName, input.Category, input.Description, auction_entity.AuctionCondition(input.Condition))

	if err := useCase.AuctionRepository.CreateAuction(ctx, auction); err != nil {
		return err
	}

	return nil
}

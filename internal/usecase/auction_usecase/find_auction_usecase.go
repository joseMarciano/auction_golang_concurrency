package auction_usecase

import (
	"auction_golang_concurrency/internal/entity/auction_entity"
	"auction_golang_concurrency/internal/internal_error"
	"auction_golang_concurrency/internal/usecase/bid_usecase"
	"context"
)

func (useCase *AuctionUseCase) FindAuctionByID(ctx context.Context, id string) (*AuctionOutputDto, *internal_error.InternalError) {
	auction, err := useCase.AuctionRepository.FindAuctionByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &AuctionOutputDto{
		Id:          auction.Id,
		ProductName: auction.ProductName,
		Category:    auction.Category,
		Description: auction.Description,
		Condition:   auction.Condition,
		Status:      auction.Status,
		CreatedAt:   auction.CreatedAt,
	}, nil
}

func (useCase *AuctionUseCase) FindAuctions(ctx context.Context, status int, category, productName string) ([]AuctionOutputDto, *internal_error.InternalError) {
	auctions, err := useCase.AuctionRepository.FindAuctions(ctx, auction_entity.AuctionStatus(status), category, productName)
	if err != nil {
		return nil, err
	}
	var output []AuctionOutputDto
	for _, auction := range auctions {
		output = append(output, AuctionOutputDto{
			Id:          auction.Id,
			ProductName: auction.ProductName,
			Category:    auction.Category,
			Description: auction.Description,
			Condition:   auction.Condition,
			Status:      auction.Status,
			CreatedAt:   auction.CreatedAt,
		})
	}
	return output, nil
}

func (useCase *AuctionUseCase) FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*WinningInfoOutputDto, *internal_error.InternalError) {
	auction, err := useCase.AuctionRepository.FindAuctionByID(ctx, auctionId)
	if err != nil {
		return nil, err
	}

	auctionOutput := &AuctionOutputDto{
		Id:          auction.Id,
		ProductName: auction.ProductName,
		Category:    auction.Category,
		Description: auction.Description,
		Condition:   auction.Condition,
		Status:      auction.Status,
		CreatedAt:   auction.CreatedAt,
	}

	bidWinning, err := useCase.BidRepository.FindWinningBidByAuctionId(ctx, auction.Id)
	if err != nil {
		return &WinningInfoOutputDto{
			Auction: auctionOutput,
			Bid:     nil,
		}, nil
	}

	bidOutput := &bid_usecase.BidOutputDto{
		Id:        bidWinning.Id,
		UserId:    bidWinning.UserId,
		AuctionId: bidWinning.AuctionId,
		Amount:    bidWinning.Amount,
		CreatedAt: bidWinning.CreatedAt,
	}

	return &WinningInfoOutputDto{
		Auction: auctionOutput,
		Bid:     bidOutput,
	}, nil
}

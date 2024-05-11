package bid_usecase

import (
	"auction_golang_concurrency/internal/internal_error"
	"context"
)

func (useCase *BidUseCase) FindByAuctionId(ctx context.Context, auctionId string) ([]BidOutputDto, *internal_error.InternalError) {
	bids, err := useCase.BidRepository.FindByAuctionId(ctx, auctionId)
	if err != nil {
		return nil, err
	}

	var outputs []BidOutputDto
	for _, bid := range bids {
		outputs = append(outputs, BidOutputDto{
			Id:        bid.Id,
			UserId:    bid.UserId,
			AuctionId: bid.AuctionId,
			Amount:    bid.Amount,
			CreatedAt: bid.CreatedAt,
		})
	}

	return outputs, nil
}

func (useCase *BidUseCase) FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*BidOutputDto, *internal_error.InternalError) {
	bid, err := useCase.BidRepository.FindWinningBidByAuctionId(ctx, auctionId)
	if err != nil {
		return nil, err
	}

	output := BidOutputDto{
		Id:        bid.Id,
		UserId:    bid.UserId,
		AuctionId: bid.AuctionId,
		Amount:    bid.Amount,
		CreatedAt: bid.CreatedAt,
	}

	return &output, nil
}

package bid_usecase

import (
	"auction_golang_concurrency/configuration/logger"
	"auction_golang_concurrency/internal/entity/bid_entity"
	"auction_golang_concurrency/internal/internal_error"
	"context"
	"os"
	"strconv"
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

	timer               *time.Timer
	maxBatchSize        int
	batchInsertInterval time.Duration
	bidChannel          chan bid_entity.Bid
}

func NewBidUseCase(bidRepository bid_entity.BidRepositoryInterface) BidUseCaseInterface {
	maxInterval := getMaxBatchInterval()
	batchSize := getMaxBatchSize()
	bidUseCase := &BidUseCase{
		BidRepository:       bidRepository,
		maxBatchSize:        batchSize,
		batchInsertInterval: maxInterval,
		timer:               time.NewTimer(maxInterval),
		bidChannel:          make(chan bid_entity.Bid, batchSize),
	}

	bidUseCase.triggerCreateRoutine(context.TODO())

	return bidUseCase
}

var bidBatch []bid_entity.Bid

type BidUseCaseInterface interface {
	CreateBid(ctx context.Context, input *BidInputDto) (*BidOutputDto, *internal_error.InternalError)
	FindByAuctionId(ctx context.Context, auctionId string) ([]BidOutputDto, *internal_error.InternalError)
	FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*BidOutputDto, *internal_error.InternalError)
}

func (useCase *BidUseCase) CreateBid(ctx context.Context, input *BidInputDto) (*BidOutputDto, *internal_error.InternalError) {
	bidEntity := bid_entity.NewBid(input.UserId, input.AuctionId, input.Amount)

	useCase.bidChannel <- *bidEntity

	return nil, nil
}

func (useCase *BidUseCase) triggerCreateRoutine(ctx context.Context) {
	go func() {
		defer close(useCase.bidChannel)

		for {
			select {
			case bidEntity, ok := <-useCase.bidChannel:
				if !ok { // se não OK. vai enviar o que tem na lista
					if len(bidBatch) > 0 {
						if err := useCase.BidRepository.CreateBid(ctx, bidBatch); err != nil {
							logger.Error("error trying to process bid batch list", err)
						}
					}
					bidBatch = nil
					return
				}

				bidBatch = append(bidBatch, bidEntity)

				if len(bidBatch) >= useCase.maxBatchSize {
					if err := useCase.BidRepository.CreateBid(ctx, bidBatch); err != nil {
						logger.Error("error trying to process bid batch list", err)
					}
					bidBatch = nil
					useCase.timer.Reset(useCase.batchInsertInterval)
				}
			case <-useCase.timer.C: // estourou o timeout
				if err := useCase.BidRepository.CreateBid(ctx, bidBatch); err != nil {
					logger.Error("error trying to process bid batch list", err)
				}
				bidBatch = nil
				useCase.timer.Reset(useCase.batchInsertInterval)
			}
		}
	}()
}

func getMaxBatchInterval() time.Duration {
	batchInsertInterval := os.Getenv("BATCH_INSERT_INTERVAL")
	duration, err := time.ParseDuration(batchInsertInterval)
	if err != nil {
		return 3 * time.Minute
	}

	return duration
}

func getMaxBatchSize() int {
	maxBatchSize := os.Getenv("MAX_BATCH_SIZE")
	size, err := strconv.Atoi(maxBatchSize)
	if err != nil {
		return 100
	}

	return size
}

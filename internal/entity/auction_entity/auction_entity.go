package auction_entity

import (
	"auction_golang_concurrency/internal/internal_error"
	"context"
	"github.com/google/uuid"
	"time"
)

func CreateAuction(productName, category, description string, condition AuctionCondition) *Auction {
	return &Auction{
		Id:          uuid.New().String(),
		ProductName: productName,
		Category:    category,
		Description: description,
		Condition:   condition,
		Status:      Active,
		CreatedAt:   time.Now(),
	}
}

type Auction struct {
	Id          string
	ProductName string
	Category    string
	Description string
	Condition   AuctionCondition
	Status      AuctionStatus
	CreatedAt   time.Time
}

type AuctionCondition int
type AuctionStatus int

const (
	Active AuctionStatus = iota + 1
	Completed
)

const (
	New AuctionCondition = iota
	Used
)

type AuctionRepositoryInterface interface {
	CreateAuction(ctx context.Context, auction *Auction) *internal_error.InternalError
	FindAuctionByID(ctx context.Context, id string) (*Auction, *internal_error.InternalError)
	FindAuctions(ctx context.Context, status AuctionStatus, category, productName string) ([]Auction, *internal_error.InternalError)
}

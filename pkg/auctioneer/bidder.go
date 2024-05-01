package auctioneer

import "errors"

var (
	ErrBidderEmptyName   = errors.New("bidder name is required")
	ErrNoInitialBid      = errors.New("initialBid is required")
	ErrNoMaxBid          = errors.New("maxBid is required")
	ErrNoIncrement       = errors.New("bidIncrement is required")
	ErrInvalidInitialBid = errors.New("startBid amount cannot be larger than maxBid")
	ErrThresholdReached  = errors.New("surpassing max bid")
)

// Bidder represents a participant in the auction.
// TODO create a currency type, no using float64
type Bidder struct {
	name         string
	initialBid   float64
	maxBid       float64
	increment    float64
	latestBid    float64
	outOfAuction bool
}

// BidderParams holds information required to build a Bidder.
type BidderParams struct {
	Name       string
	InitialBid float64
	MaxBid     float64
	Increment  float64
}

// NewBidder allocates and returns a new [Bidder].
func NewBidder(params BidderParams) (*Bidder, error) {
	if params.Name == "" {
		return nil, ErrBidderEmptyName
	}
	if params.InitialBid == 0 {
		return nil, ErrNoInitialBid
	}
	if params.MaxBid == 0 {
		return nil, ErrNoMaxBid
	}
	if params.Increment == 0 {
		return nil, ErrNoIncrement
	}
	if params.InitialBid > params.MaxBid {
		return nil, ErrInvalidInitialBid
	}

	return &Bidder{
		name:       params.Name,
		initialBid: params.InitialBid,
		maxBid:     params.MaxBid,
		increment:  params.Increment,
		latestBid:  params.InitialBid,
	}, nil
}

// IncrementBid increases the latestBid based on the configured increment amount.
func (b *Bidder) IncrementBid() error {
	newBid := b.latestBid + b.increment
	if newBid > b.maxBid {
		return ErrThresholdReached
	}
	b.latestBid = newBid
	return nil
}

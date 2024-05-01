package auctioneer

import (
	"errors"
	"math"
)

var (
	ErrNoContest = errors.New("not enough participants. need at least 2")
)

// Auctioneer manages the auction.
type Auctioneer struct {
	maxRounds int
	// ... more fields like max bidders, etc.
}

// default number of max rounds to avoid
// never ending or high number of iterations
const defaultMaxRounds = 30

func defaultAuctioneer() *Auctioneer {
	return &Auctioneer{
		maxRounds: defaultMaxRounds,
	}
}

// Option is a functional option for configuring the Auctioneer
type Option func(*Auctioneer)

// NewAuctioner allocates and returns a new [Auctioneer].
func NewAuctioneer(opts ...Option) *Auctioneer {
	auctioneer := defaultAuctioneer()
	for _, opt := range opts {
		opt(auctioneer)
	}
	return auctioneer
}

// WithMaxRounds sets the auctioneer max rounds for the auction.
func WithMaxRounds(rounds int) Option {
	return func(a *Auctioneer) {
		a.maxRounds = rounds
	}
}

// Auction mimics an auction based on multiple rounds.
func (a *Auctioneer) Auction(bidders []*Bidder) (*Bidder, error) {
	if len(bidders) < 2 {
		return nil, ErrNoContest
	}

	remainingPlayers := len(bidders)
	winningBidder := &Bidder{}
	var winningBid float64
	for i := 1; i <= a.maxRounds; i++ {
		if remainingPlayers == 0 {
			break
		}

		for _, b := range bidders {
			if b.outOfAuction {
				continue
			}

			if b.latestBid < winningBid {
				err := b.IncrementBid()
				if err != nil {
					b.outOfAuction = true
					remainingPlayers--
					continue
				}
			}

			if b.latestBid > winningBid {
				winningBid = b.latestBid
				winningBidder = b
				if remainingPlayers == 1 {
					break
				}
			}
		}
	}

	return winningBidder, nil
}

// Alternative approach to calculate max possible bid.
// This solution could be considered more efficient,
// doesn't cover the rule of lowest winning bid.
func (a *Auctioneer) AuctionAlt(bidders []*Bidder) (*Bidder, error) {
	if len(bidders) < 2 {
		return nil, ErrNoContest
	}

	var winningBidder *Bidder
	var winningBid float64
	for _, b := range bidders {
		numIncrements := math.Floor((b.maxBid - b.initialBid) / b.increment)
		bidAtMax := b.initialBid + (numIncrements * b.increment)
		if bidAtMax > winningBid {
			winningBid = bidAtMax
			winningBidder = b
		}
	}

	return winningBidder, nil
}

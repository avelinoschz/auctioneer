package main

import (
	"log"

	"github.com/avelinoschz/auctioneer/pkg/auctioneer"
)

func main() {
	a := auctioneer.NewAuctioneer(
		auctioneer.WithMaxRounds(99),
	)

	mike, err := auctioneer.NewBidder(auctioneer.BidderParams{
		Name:       "mike",
		InitialBid: 2500,
		MaxBid:     3000,
		Increment:  500,
	})
	if err != nil {
		log.Fatal(err)
	}

	avelino, _ := auctioneer.NewBidder(auctioneer.BidderParams{
		Name:       "avelino",
		InitialBid: 2800,
		MaxBid:     3100,
		Increment:  201,
	})
	if err != nil {
		log.Fatal(err)
	}

	jeff, _ := auctioneer.NewBidder(auctioneer.BidderParams{
		Name:       "jeff",
		InitialBid: 2501,
		MaxBid:     3200,
		Increment:  247,
	})
	if err != nil {
		log.Fatal(err)
	}

	bidders := []*auctioneer.Bidder{mike, avelino, jeff}
	winner, err := a.Auction(bidders)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("winner is %+v\n", winner)
}

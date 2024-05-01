package auctioneer_test

import (
	"testing"

	"github.com/avelinoschz/auctioneer/pkg/auctioneer"
)

func TestAuction(t *testing.T) {
	t.Parallel()

	sasha, _ := auctioneer.NewBidder(auctioneer.BidderParams{
		Name:       "sasha",
		InitialBid: 50,
		MaxBid:     80,
		Increment:  3,
	})
	john, _ := auctioneer.NewBidder(auctioneer.BidderParams{
		Name:       "john",
		InitialBid: 60,
		MaxBid:     82,
		Increment:  2,
	})
	pat, _ := auctioneer.NewBidder(auctioneer.BidderParams{
		Name:       "pat",
		InitialBid: 55,
		MaxBid:     85,
		Increment:  5,
	})

	riley, _ := auctioneer.NewBidder(auctioneer.BidderParams{
		Name:       "riley",
		InitialBid: 700,
		MaxBid:     725,
		Increment:  2,
	})
	morgan, _ := auctioneer.NewBidder(auctioneer.BidderParams{
		Name:       "morgan",
		InitialBid: 599,
		MaxBid:     725,
		Increment:  15,
	})
	charlie, _ := auctioneer.NewBidder(auctioneer.BidderParams{
		Name:       "charlie",
		InitialBid: 625,
		MaxBid:     725,
		Increment:  8,
	})

	alex, _ := auctioneer.NewBidder(auctioneer.BidderParams{
		Name:       "alex",
		InitialBid: 2500,
		MaxBid:     3000,
		Increment:  500,
	})
	jesse, _ := auctioneer.NewBidder(auctioneer.BidderParams{
		Name:       "jesse",
		InitialBid: 2800,
		MaxBid:     3100,
		Increment:  201,
	})
	drew, _ := auctioneer.NewBidder(auctioneer.BidderParams{
		Name:       "drew",
		InitialBid: 2501,
		MaxBid:     3200,
		Increment:  247,
	})

	testCases := []struct {
		desc  string
		input []*auctioneer.Bidder
		want  *auctioneer.Bidder
		err   error
	}{
		{
			desc:  "no contest - no participants",
			input: []*auctioneer.Bidder{},
			err:   auctioneer.ErrNoContest,
		},
		{
			desc: "no contest - one participant",
			input: []*auctioneer.Bidder{
				sasha,
			},
			err: auctioneer.ErrNoContest,
		},
		{
			desc: "auction 1",
			input: []*auctioneer.Bidder{
				sasha, john, pat,
			},
			want: pat,
		},
		{
			desc: "auction 2",
			input: []*auctioneer.Bidder{
				riley, morgan, charlie,
			},
			want: riley,
		},
		{
			desc: "auction 3",
			input: []*auctioneer.Bidder{
				alex, jesse, drew,
			},
			want: jesse,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			auctioneer := auctioneer.NewAuctioneer(auctioneer.WithMaxRounds(99))
			got, err := auctioneer.Auction(tc.input)
			if err != nil && tc.err.Error() != err.Error() {
				t.Fatalf("error doesn't match. got: %v, want: %v", err, tc.err)
			}
			if tc.want != got {
				t.Fatalf("wrong winner. got: %+v, want: %+v", got, tc.want)
			}
		})
	}
}

func TestAuctionAlt(t *testing.T) {
	t.Parallel()

	sasha, _ := auctioneer.NewBidder(auctioneer.BidderParams{
		Name:       "sasha",
		InitialBid: 50,
		MaxBid:     80,
		Increment:  3,
	})
	john, _ := auctioneer.NewBidder(auctioneer.BidderParams{
		Name:       "john",
		InitialBid: 60,
		MaxBid:     82,
		Increment:  2,
	})
	pat, _ := auctioneer.NewBidder(auctioneer.BidderParams{
		Name:       "pat",
		InitialBid: 55,
		MaxBid:     85,
		Increment:  5,
	})

	riley, _ := auctioneer.NewBidder(auctioneer.BidderParams{
		Name:       "riley",
		InitialBid: 700,
		MaxBid:     725,
		Increment:  2,
	})
	morgan, _ := auctioneer.NewBidder(auctioneer.BidderParams{
		Name:       "morgan",
		InitialBid: 599,
		MaxBid:     725,
		Increment:  15,
	})
	charlie, _ := auctioneer.NewBidder(auctioneer.BidderParams{
		Name:       "charlie",
		InitialBid: 625,
		MaxBid:     725,
		Increment:  8,
	})

	alex, _ := auctioneer.NewBidder(auctioneer.BidderParams{
		Name:       "alex",
		InitialBid: 2500,
		MaxBid:     3000,
		Increment:  500,
	})
	jesse, _ := auctioneer.NewBidder(auctioneer.BidderParams{
		Name:       "jesse",
		InitialBid: 2800,
		MaxBid:     3100,
		Increment:  201,
	})
	drew, _ := auctioneer.NewBidder(auctioneer.BidderParams{
		Name:       "drew",
		InitialBid: 2501,
		MaxBid:     3200,
		Increment:  247,
	})

	testCases := []struct {
		desc   string
		input  []*auctioneer.Bidder
		want   *auctioneer.Bidder
		hasErr bool
	}{
		{
			desc: "alt auction 1",
			input: []*auctioneer.Bidder{
				sasha, john, pat,
			},
			want: pat,
		},
		{
			desc: "alt auction 2",
			input: []*auctioneer.Bidder{
				riley, morgan, charlie,
			},
			want: riley,
		},
		{
			desc: "alt auction 3",
			input: []*auctioneer.Bidder{
				alex, jesse, drew,
			},
			want: jesse,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			auctioneer := auctioneer.NewAuctioneer()
			got, err := auctioneer.AuctionAlt(tc.input)
			if (err != nil) != tc.hasErr {
				t.Fatalf("error doesn't match. got: %v, want: %v", err, tc)
			}
			if tc.want != got {
				t.Fatalf("wrong winner. got: %+v, want: %+v", got, tc.want)
			}
		})
	}
}

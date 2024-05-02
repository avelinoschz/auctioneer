package auctioneer

import "testing"

func TestNewBidder(t *testing.T) {
	t.Parallel()

	bidderName := "bidderName"
	initialBid := 100.00
	maxBid := 200.00
	increment := 10.00
	wrongInitialBid := 999.00

	testCases := []struct {
		desc  string
		input BidderParams
		err   error
	}{
		{
			desc:  "empty name",
			input: BidderParams{},
			err:   ErrBidderEmptyName,
		},
		{
			desc: "no initial bid",
			input: BidderParams{
				Name: bidderName,
			},
			err: ErrNoInitialBid,
		},
		{
			desc: "no max bid",
			input: BidderParams{
				Name:       bidderName,
				InitialBid: initialBid,
			},
			err: ErrNoMaxBid,
		},
		{
			desc: "no increment",
			input: BidderParams{
				Name:       bidderName,
				InitialBid: initialBid,
				MaxBid:     maxBid,
			},
			err: ErrNoIncrement,
		},
		{
			desc: "initial bid bigger than max bid",
			input: BidderParams{
				Name:       bidderName,
				InitialBid: wrongInitialBid,
				MaxBid:     maxBid,
				Increment:  increment,
			},
			err: ErrInvalidInitialBid,
		},
		{
			desc: "base case",
			input: BidderParams{
				Name:       bidderName,
				InitialBid: initialBid,
				MaxBid:     maxBid,
				Increment:  increment,
			},
			err: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			_, err := NewBidder(tc.input)
			if err != nil && tc.err.Error() != err.Error() {
				t.Fatalf("error doesn't match. got: %v, want: %v", err, tc.err)
			}
		})
	}
}

func TestIncrementBid(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc  string
		input *Bidder
		want  float64
		err   error
	}{
		{
			desc: "base case",
			input: &Bidder{
				latestBid: 1000,
				maxBid:    1500,
				increment: 200,
			},
			want: 1200,
			err:  nil,
		},
		{
			desc: "threshold reached",
			input: &Bidder{
				latestBid: 1000,
				maxBid:    1500,
				increment: 550,
			},
			err: ErrThresholdReached,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			bidder := tc.input
			err := bidder.incrementBid()
			if err != nil && tc.err.Error() != err.Error() {
				t.Fatalf("error doesn't match. got: %v, want: %v", err, tc.err)
			}
			if tc.err != nil {
				return
			}

			if tc.want != bidder.latestBid {
				t.Fatalf("wrong latest bid. got: %+v, want: %+v", bidder.latestBid, tc.want)
			}
		})
	}
}

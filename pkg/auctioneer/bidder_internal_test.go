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
		want  *Bidder
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
			want: &Bidder{
				name:       bidderName,
				initialBid: initialBid,
				maxBid:     maxBid,
				increment:  increment,
				latestBid:  initialBid,
			},
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

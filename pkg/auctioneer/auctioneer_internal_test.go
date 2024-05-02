package auctioneer

import "testing"

func TestNewAuctioneer(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc  string
		input Option
		want  Auctioneer
	}{
		{
			desc:  "default opts",
			input: func(a *Auctioneer) {},
			want: Auctioneer{
				maxRounds: 30,
			},
		},
		{
			desc:  "with max rounds",
			input: WithMaxRounds(99),
			want: Auctioneer{
				maxRounds: 99,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got := NewAuctioneer(tc.input)
			if tc.want != *got {
				t.Fatalf("wrong configuration. got: %+v, want: %+v", got, tc.want)
			}
		})
	}
}

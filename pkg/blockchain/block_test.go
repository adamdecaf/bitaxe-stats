package blockchain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculateDifficulty(t *testing.T) {
	cases := []struct {
		hash      string
		formatted string
	}{
		{
			hash:      "0000000000000000000000000000000000000000000000000000000000000000",
			formatted: "0",
		},
		{
			hash:      "000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f",
			formatted: "2.54K",
		},
		{
			hash:      "000000000003ba27aa200b1cecaad478d2b00432346c3f1f3986da1afd33e506",
			formatted: "17.58K",
		},
		{
			hash:      "00000000000000000001b7fd44963a725914a8db0998145d28615f32065d3d9c",
			formatted: "163.77T",
		},
		{
			hash:      "0000000000000000000000005d6f06154c8685146aa7bc3dc9843876c9cefd0f",
			formatted: "50.54E",
		},
		{
			hash:      "000000000000000000006414aea39be567cf1d5ff6cbf2d77254fe7c714b0d81",
			formatted: "719.98T",
		},
		{
			hash:      "000000000000000000022247ce3898fc9e10be0161cde5f369293d7a916ca8b3",
			formatted: "131.9T",
		},
	}
	for _, tc := range cases {
		t.Run(tc.hash, func(t *testing.T) {
			diff, err := CalculateDifficulty(tc.hash)
			require.NoError(t, err)

			v, _ := diff.Float64()
			d := Difficulty{RawValue: v}

			require.Equal(t, tc.formatted, d.Format())
		})
	}
}

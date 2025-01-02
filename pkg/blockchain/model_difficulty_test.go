package blockchain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseDifficulty(t *testing.T) {
	cases := []struct {
		input    string
		expected Difficulty
		error    error
	}{
		{
			input:    "10",
			expected: Difficulty{RawValue: 10, Unit: ""},
		},
		{
			input:    "1 k",
			expected: Difficulty{RawValue: 1000, Unit: "K"},
		},
		{
			input:    "1.2 k",
			expected: Difficulty{RawValue: 1200, Unit: "K"},
		},
		{
			input:    "1.2M",
			expected: Difficulty{RawValue: 1_200_000, Unit: "M"},
		},
		{
			input:    "115M",
			expected: Difficulty{RawValue: 115_000_000, Unit: "M"},
		},
		{
			input:    "6.97G",
			expected: Difficulty{RawValue: 6_970_000_000, Unit: "G"},
		},
		{
			input:    "3.29T",
			expected: Difficulty{RawValue: 3_290_000_000_000, Unit: "T"},
		},
		{
			input:    "42.859P",
			expected: Difficulty{RawValue: 42_859_000_000_000_000, Unit: "P"},
		},
		{
			input:    "5.1E",
			expected: Difficulty{RawValue: 5_100_000_000_000_000_000, Unit: "E"},
		},
		{
			input:    "432.1236897E",
			expected: Difficulty{RawValue: 432_123_689_700_000_000_000, Unit: "E"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			diff, err := ParseDifficulty(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, diff)
		})
	}
}

package blockchain

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseDifficulty(t *testing.T) {
	cases := []struct {
		input     string
		expected  Difficulty
		formatted string // force a manual comparison
		error     error
	}{
		{
			input:     "10",
			expected:  Difficulty{RawValue: 10, Unit: ""},
			formatted: "10",
		},
		{
			input:     "1 k",
			expected:  Difficulty{RawValue: 1000, Unit: "K"},
			formatted: "1K",
		},
		{
			input:     "1.2 k",
			expected:  Difficulty{RawValue: 1200, Unit: "K"},
			formatted: "1.2K",
		},
		{
			input:     "1.2M",
			expected:  Difficulty{RawValue: 1_200_000, Unit: "M"},
			formatted: "1.2M",
		},
		{
			input:     "115M",
			expected:  Difficulty{RawValue: 115_000_000, Unit: "M"},
			formatted: "115M",
		},
		{
			input:     "6.97G",
			expected:  Difficulty{RawValue: 6_970_000_000, Unit: "G"},
			formatted: "6.97G",
		},
		{
			input:     "3.29T",
			expected:  Difficulty{RawValue: 3_290_000_000_000, Unit: "T"},
			formatted: "3.29T",
		},
		{
			input:     "42.859P",
			expected:  Difficulty{RawValue: 42_859_000_000_000_000, Unit: "P"},
			formatted: "42.86P",
		},
		{
			input:     "5.1E",
			expected:  Difficulty{RawValue: 5_100_000_000_000_000_000, Unit: "E"},
			formatted: "5.1E",
		},
		{
			input:     "432.1236897E",
			expected:  Difficulty{RawValue: 432_123_689_700_000_000_000, Unit: "E"},
			formatted: "432.12E",
		},
		{
			input:     "13912524048946",
			expected:  Difficulty{RawValue: 13_912_524_048_946, Unit: ""},
			formatted: "13.91T",
		},
		{
			input:     "13732352106018.34",
			expected:  Difficulty{RawValue: 13_732_352_106_018, Unit: ""},
			formatted: "13.73T",
		},
		{
			input:     "113757508810854",
			expected:  Difficulty{RawValue: 113_757_508_810_854, Unit: ""},
			formatted: "113.76T",
		},
		{
			input:     "187842912625266",
			expected:  Difficulty{RawValue: 187_842_912_625_266, Unit: ""},
			formatted: "187.84T",
		},
		{
			input:     "719983272352941",
			expected:  Difficulty{RawValue: 719_983_272_352_941, Unit: ""},
			formatted: "719.98T",
		},
		{
			input:     "31360548173144",
			expected:  Difficulty{RawValue: 31_360_548_173_144, Unit: ""},
			formatted: "31.36T",
		},
		{
			input:     "541668626379227136",
			expected:  Difficulty{RawValue: 541_668_626_379_227_136, Unit: ""},
			formatted: "541.67P",
		},
		{
			input:     "50541668626379227136",
			expected:  Difficulty{RawValue: 50_541_668_626_379_227_136, Unit: ""},
			formatted: "50.54E",
		},
		{
			input:     "32150541668626380881920",
			expected:  Difficulty{RawValue: 32_150_541_668_626_379_227_136, Unit: ""},
			formatted: "32.15Z",
		},
		// {
		// 	input:     "87232150541668626379227136",
		// 	expected:  Difficulty{RawValue: 87_232_150_541_668_632_923_996_160, Unit: ""},
		// 	formatted: "87.23Y",
		// },
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			diff, err := ParseDifficulty(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, diff)

			// Test Format()
			require.Equal(t, tc.formatted, diff.Format())

			// Test the String() method
			cleanInput := strings.ToUpper(strings.ReplaceAll(tc.input, " ", ""))

			// For values without units, the String() should return the RawValue without decimal places
			if diff.Unit == "" {
				// For decimal values without units, String() won't preserve decimals
				if strings.Contains(tc.input, ".") {
					// Take only the integer part
					cleanInput = strings.Split(cleanInput, ".")[0]
				}
			}

			require.Equal(t, cleanInput, diff.String())
		})
	}
}

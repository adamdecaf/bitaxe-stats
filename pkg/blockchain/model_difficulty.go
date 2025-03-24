package blockchain

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Difficulty struct {
	RawValue float64
	Unit     string
}

func (d Difficulty) String() string {
	// For raw values without units, just return the raw number formatted without decimals
	if d.Unit == "" {
		return fmt.Sprintf("%.0f", d.RawValue)
	}

	// For values with units, use the correct divisor
	value := d.RawValue
	divisor := 1.0
	for i := 1; ; i++ {
		power := math.Pow(10, float64(3*i)) // Calculate powers of 1000 (10^3, 10^6, etc.)
		if value < power {
			break
		}
		divisor = power
	}

	// Format without trailing zeros
	num := strconv.FormatFloat(value/divisor, 'f', -1, 64)
	return fmt.Sprintf("%s%s", num, d.Unit)
}

// Format returns a human-readable representation of difficulty with appropriate units
func (d Difficulty) Format() string {
	// Get the raw value
	value := d.RawValue

	// Units defined by the Bitcoin specification (and beyond)
	// K = kilo (10^3), M = mega (10^6), G = giga (10^9), T = tera (10^12)
	// P = peta (10^15), E = exa (10^18), Z = zetta (10^21), Y = yotta (10^24)
	units := []string{"", "K", "M", "G", "T", "P", "E", "Z", "Y"}
	unitIndex := 0

	for i := 0; i < len(units)-1; i++ {
		if value < math.Pow(10, float64(3*(i+1))) {
			break
		}
		unitIndex++
	}

	// Calculate the divided value
	divisor := math.Pow(10, float64(3*unitIndex))
	scaledValue := value / divisor

	// Format the output - if it's a whole number, show no decimal places
	if scaledValue == math.Floor(scaledValue) {
		return fmt.Sprintf("%.0f%s", scaledValue, units[unitIndex])
	}

	// Otherwise show with up to 2 decimal places, no trailing zeros
	formatted := strconv.FormatFloat(scaledValue, 'f', 2, 64)
	formatted = strings.TrimRight(formatted, "0")
	formatted = strings.TrimRight(formatted, ".")

	return fmt.Sprintf("%s%s", formatted, units[unitIndex])
}

func ParseDifficulty(raw string) (Difficulty, error) {
	var diff Difficulty

	raw = strings.TrimSpace(raw)
	if raw == "" {
		return diff, errors.New("blank difficulty input")
	}

	// Check if the last character is a letter (unit)
	lastChar := raw[len(raw)-1:]
	isUnit := false
	if lastChar >= "a" && lastChar <= "z" || lastChar >= "A" && lastChar <= "Z" {
		isUnit = true
	}

	// Parse the value and unit
	var valueStr string
	if isUnit {
		// Extract unit
		diff.Unit = strings.ToUpper(lastChar)
		valueStr = strings.TrimSpace(raw[:len(raw)-1])
	} else {
		// No unit
		diff.Unit = ""
		valueStr = raw
	}

	// Parse value
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return diff, fmt.Errorf("parsing difficulty: %w", err)
	}

	// Scale based on unit
	switch diff.Unit {
	case "Y": // yotta (10^24)
		value *= 1_000
		fallthrough
	case "Z": // zetta (10^21)
		value *= 1_000
		fallthrough
	case "E": // exa (10^18)
		value *= 1_000
		fallthrough
	case "P": // peta (10^15)
		value *= 1_000
		fallthrough
	case "T": // tera (10^12)
		value *= 1_000
		fallthrough
	case "G": // giga (10^9)
		value *= 1_000
		fallthrough
	case "M": // mega (10^6)
		value *= 1_000
		fallthrough
	case "K": // kilo (10^3)
		value *= 1_000
	default:
		// No unit, truncate decimal for raw values
		if strings.Contains(raw, ".") {
			value = math.Floor(value)
		}
	}

	diff.RawValue = value
	return diff, nil
}

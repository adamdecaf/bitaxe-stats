package blockchain

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Difficulty struct {
	RawValue float64
	Unit     string
}

func (d Difficulty) String() string {
	return fmt.Sprintf("%.2f%s", d.RawValue, d.Unit)
}

func ParseDifficulty(raw string) (Difficulty, error) {
	var diff Difficulty

	raw = strings.TrimSpace(raw)
	if raw == "" {
		return diff, errors.New("blank difficulty input")
	}

	// Grab the unit
	unit := raw[len(raw)-1:]

	// Grab the value
	v := strings.TrimSpace(strings.TrimSuffix(raw, unit))
	value, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return diff, fmt.Errorf("parsing difficulty: %w", err)
	}

	diff.Unit = strings.ToUpper(unit)

	// Scale
	// TODO(adam): if value > 100 reduce scale by
	switch diff.Unit {
	case "E":
		value *= 1_000
		fallthrough
	case "P":
		value *= 1_000
		fallthrough
	case "T":
		value *= 1_000
		fallthrough
	case "G":
		value *= 1_000
		fallthrough
	case "M":
		value *= 1_000
		fallthrough
	case "K":
		value *= 1_000
	default:
		// no unit, reparse raw as RawValue
		diff.Unit = "" // no units

		value, err = strconv.ParseFloat(raw, 64)
		if err != nil {
			return diff, fmt.Errorf("reparsing without units: %w", err)
		}
	}
	diff.RawValue = value

	return diff, nil
}

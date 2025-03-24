package blockchain

import (
	"encoding/hex"
	"fmt"
	"math/big"
)

func HashDifficulty(hash string) (Difficulty, error) {
	diff, err := CalculateDifficulty(hash)
	if err != nil {
		return Difficulty{}, err
	}

	v, _ := diff.Float64()

	return Difficulty{RawValue: v}, nil
}

// CalculateDifficulty calculates the difficulty of a block based on its hash
func CalculateDifficulty(blockHash string) (*big.Float, error) {
	hashBytes, err := hex.DecodeString(blockHash)
	if err != nil {
		return nil, fmt.Errorf("invalid hash format: %v", err)
	}

	hashInt := new(big.Int).SetBytes(hashBytes)

	// In Bitcoin, the difficulty calculation uses a reference target
	// This is the "difficulty 1" target from Bitcoin (0x00000000FFFF0000000000000000000000000000000000000000000000000000)
	// Using this as a reference produces the expected Bitcoin difficulty values
	refTarget := new(big.Int)
	refTarget.SetString("00000000FFFF0000000000000000000000000000000000000000000000000000", 16)

	// Calculate difficulty: refTarget / hashInt (standard Bitcoin formula)
	// The bigger the hashInt, the smaller the difficulty
	if hashInt.Cmp(big.NewInt(0)) == 0 {
		return big.NewFloat(0.0), nil
	}

	difficulty := new(big.Float).SetInt(refTarget)
	hashFloat := new(big.Float).SetInt(hashInt)
	difficulty = difficulty.Quo(difficulty, hashFloat)

	return difficulty, nil
}

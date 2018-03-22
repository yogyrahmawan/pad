package pad

import (
	"math"
	"strings"
)

func repeat(lenInput, lenPad, padLength int) int {
	return int(math.Ceil(float64(padLength-lenInput)/float64(lenPad) + float64(1)))
}

// Right implements padding right
func Right(inputStr, padStr string, padLength int) string {
	if len(inputStr) > padLength {
		return inputStr
	}

	return (inputStr + strings.Repeat(padStr, repeat(len(inputStr), len(padStr), padLength)))[:padLength]
}

// Left implements padding left
func Left(inputStr, padStr string, padLength int) string {
	if len(inputStr) > padLength {
		return inputStr
	}

	out := strings.Repeat(padStr, repeat(len(inputStr), len(padStr), padLength)) + inputStr
	return out[len(out)-padLength:]
}

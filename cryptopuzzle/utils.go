package cryptopuzzle

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"math/bits"
	"time"
)

// generateRandomBytes generates random bytes
// source: https://gist.github.com/6220119/7ca4244528ac65abef3a39c8a2ec7ea3
func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// calculateAuthenticator, using the secret, calculates an authenticator for a given puzzle
func calculateAuthenticator(secret []byte, puzzle Puzzle) (Puzzle, error) {

	h := sha256.New()

	_, err := fmt.Fprintf(h, "%s%d%d%d%s", puzzle.RandomBytes, puzzle.Difficulty, puzzle.Time, puzzle.Deadline, secret)
	if err != nil {
		return puzzle, err
	}

	puzzle.Authenticator = h.Sum(nil)

	return puzzle, nil
}

// isPuzzleAuthentic checks the authenticty of a puzzle
func isPuzzleAuthentic(secret []byte, puzzle Puzzle) (bool, error) {

	trustedPuzzle, err := calculateAuthenticator(secret, puzzle)
	if err != nil {
		return false, err
	}

	return bytes.Equal(puzzle.Authenticator, trustedPuzzle.Authenticator), nil
}

func isDeadlinePassed(puzzle Puzzle) bool {

	deadline := time.UnixMilli(puzzle.Time).Add(time.Millisecond * time.Duration(puzzle.Deadline))
	currentTime := time.Now()

	return deadline.Before(currentTime)
}

func isSolutionValid(puzzle Puzzle) (bool, error) {

	h := sha256.New()

	_, err := fmt.Fprintf(h, "%s%d", base64.StdEncoding.EncodeToString(puzzle.RandomBytes), puzzle.Solution)
	if err != nil {
		return false, err
	}

	solution := h.Sum(nil)

	leadingZeroCount := countLeadingZeros(solution)

	log.Printf("leading zero count %d\n", leadingZeroCount)
	result := leadingZeroCount >= int(puzzle.Difficulty)

	return result, nil
}

func countLeadingZeros(solution []byte) int {

	totalZeroCount := 0
	for i := 0; i < len(solution); i++ {

		zeroCount := bits.LeadingZeros8(solution[i])
		totalZeroCount += zeroCount
		if zeroCount != 8 {
			return totalZeroCount
		}

	}

	return totalZeroCount
}

func solvePuzzle(puzzle Puzzle) (Puzzle, error) {

	h := sha256.New()
	var err error
	var i uint64

	// JSON encode encodes data using base64. This makes things on the client side easier
	encodedString := base64.StdEncoding.EncodeToString(puzzle.RandomBytes)

	for i = 0; ; i++ {

		_, err = fmt.Fprintf(h, "%s%d", encodedString, i)
		if err != nil {
			return puzzle, err
		}

		solution := h.Sum(nil)

		if countLeadingZeros(solution) >= int(puzzle.Difficulty) {
			puzzle.Solution = uint64(i)
			return puzzle, nil
		}

		h.Reset()
	}
}

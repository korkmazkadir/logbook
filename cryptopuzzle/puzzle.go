package cryptopuzzle

import (
	"fmt"
	"time"
)

var (
	ErrPuzzleNotAuthentic = fmt.Errorf("puzzle is not authentic")
	ErrDeadlineIsPassed   = fmt.Errorf("puzzle is not solved before the defined deadline")
)

// secret is byte string that is used to create authenticators for puzzles
// it must be kept secret
var secret []byte

func init() {
	var err error
	// creates a random 32 bytes to use as secret to authanticate puzzles
	secret, err = generateRandomBytes(32)

	if err != nil {
		panic(err)
	}
}

type Puzzle struct {
	RandomBytes []byte
	Difficulty  int

	Time     int64
	Deadline int64

	Authenticator []byte

	Solution uint64
}

func CreatePuzzle(difficulty int, deadline time.Duration) (Puzzle, error) {

	randomBytes, err := generateRandomBytes(32)
	if err != nil {
		return Puzzle{}, err
	}

	p := Puzzle{RandomBytes: randomBytes, Difficulty: difficulty, Time: time.Now().UnixMilli(), Deadline: deadline.Milliseconds()}
	p, err = calculateAuthenticator(secret, p)
	if err != nil {
		return Puzzle{}, err
	}

	return p, nil
}

func VerifyPuzzleSolution(puzzle Puzzle) (bool, error) {

	isAuthentic, err := isPuzzleAuthentic(secret, puzzle)
	if err != nil {
		return false, err
	}

	if !isAuthentic {
		return false, ErrPuzzleNotAuthentic
	}

	if isDeadlinePassed(puzzle) {
		return false, ErrDeadlineIsPassed
	}

	return isSolutionValid(puzzle)
}

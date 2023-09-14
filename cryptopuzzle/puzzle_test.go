package cryptopuzzle

import (
	"testing"
	"time"
)

func TestVerifyPuzzleSolution(t *testing.T) {

	var Deadline = 31 * time.Second

	puzzle, err := CreatePuzzle(20, Deadline)

	if err != nil {
		t.Errorf("could not create puzzle: %s", err.Error())
	}

	start := time.Now()
	puzzle, err = solvePuzzle(puzzle)
	elapsed := time.Since(start)

	t.Logf("puzzle solution found in %s\n", elapsed)

	if err != nil {
		t.Errorf("could not solve puzzle: %s", err.Error())
	}

	t.Logf("puzzle solution: %d", puzzle.Solution)

	var result bool
	result, err = VerifyPuzzleSolution(puzzle)

	if err != nil {
		t.Errorf("error occured during puzzle verification: %s", err.Error())
	}

	if !result {
		t.Error("puzzle is not valid")
	}

	puzzle.Difficulty += 1

	_, err = VerifyPuzzleSolution(puzzle)

	if err != ErrPuzzleNotAuthentic {
		t.Errorf("expected error %s got %s", ErrPuzzleNotAuthentic.Error(), err.Error())
	}

	var puzzle2 Puzzle
	puzzle2, err = CreatePuzzle(12, Deadline)
	if err != nil {
		t.Errorf("could not create puzzle: %s", err.Error())
	}

	puzzle2.Authenticator[0]++

	_, err = VerifyPuzzleSolution(puzzle)
	if err != ErrPuzzleNotAuthentic {
		t.Errorf("expected error %s got %s", ErrPuzzleNotAuthentic.Error(), err.Error())
	}

}

func TestPuzzleDeadline(t *testing.T) {

	var Deadline = 500 * time.Millisecond

	puzzle, err := CreatePuzzle(33, Deadline)

	if err != nil {
		t.Errorf("could not create puzzle: %s", err.Error())
	}

	time.Sleep(600 * time.Millisecond)

	_, err = VerifyPuzzleSolution(puzzle)

	if err != ErrDeadlineIsPassed {
		t.Errorf("expected error %s got %s", ErrDeadlineIsPassed.Error(), err.Error())
	}

}

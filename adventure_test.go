package adventure

import (
	"io"
	"strings"
	"testing"
)

func TestAdventure(t *testing.T) {
	in := "5 12\n0 1\n0 2\n1 0\n1 3\n2 0\n2 3\n2 4\n3 1\n3 2\n3 4\n4 2\n4 3\n0 4\n"
	var r io.Reader = strings.NewReader(in)

	adv, _ := Read(r)

	stepsScript := strings.NewReader("1\n4\n3\n0\n2\n4\n")
	moves := []struct {
		isErr  bool
		isExit bool
		curPos int
	}{
		{false, false, 1},
		{true, false, 1},
		{false, false, 3},
		{true, false, 3},
		{false, false, 2},
		{false, true, 4},
	}

	for i, move := range moves {
		err := adv.Move(stepsScript)

		if move.curPos != adv.playerPos {
			t.Fatalf("%d: expected current position %d, got %d", i, move.curPos, adv.playerPos)
		}
		if move.isErr != (err != nil) {
			t.Fatalf("%d: expected move is error", i)
		}

		if move.isExit != adv.IsExit() {
			t.Fatalf("%d: expected exit", i)
		}
	}
}

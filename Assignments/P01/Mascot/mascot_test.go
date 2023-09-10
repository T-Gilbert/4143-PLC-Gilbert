package mascot_test

import (
	"testing"

	"github.com/T-Gilbert/4143-PLC-Gilbert/tree/main/Assignments/P01/mascot"
)

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Go Gopher" {
		t.Fatal("Wrong mascot :(")
	}
}

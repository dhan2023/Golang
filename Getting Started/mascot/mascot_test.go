package mascot_test

import (
	"testing"

	"github.com/dhan2023/mascot"
)

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Go Gopher" {
		t.Fatal("wrong mascot :(")
	}
}

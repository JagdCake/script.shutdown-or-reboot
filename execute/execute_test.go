package execute

import (
	"testing"
)

func TestCommandOutput(t *testing.T) {
	expected := "test"
	got, _ := Command("printf", "test")

	if expected != got.String() {
		t.Errorf("Expected: %s\nGot: %s", expected, got)
	}
}

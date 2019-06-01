package menu

import "testing"

func TestMenuCreation(t *testing.T) {
	menuOptions := []string{
		"Open",
		"Close",
		"Quit",
	}

	var expected string = "1) Open\n2) Close\n3) Quit\n# "
	var got string = CreateWith(menuOptions)

	if expected != got {
		t.Errorf("Expected: %s\nGot: %s", expected, got)
	}
}

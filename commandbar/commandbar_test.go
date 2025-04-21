package commandbar

import (
	"testing"
)

func TestView(t *testing.T) {
	command := CommandMap{"q": "quit"}
	commandbar := NewCommandBar(command)
	want := " q:quit\t"

	got := commandbar.View()

	if got != want {
		t.Errorf("Wanted \"%s\", got \"%s\"", want, got)
	}
}

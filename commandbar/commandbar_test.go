package commandbar

import (
	"testing"
)

func Test_Single_Command_View(t *testing.T) {
	command := CommandMap{"q": "quit"}
	commandbar := NewCommandBar(command)
	want := " q:quit\t"

	got := commandbar.View()

	if got != want {
		t.Errorf("Wanted \"%s\", got \"%s\"", want, got)
	}
}

func Test_Multi_Command_View(t *testing.T) {
	command := CommandMap{"q": "quit", "j": "down", "k": "up"}
	commandbar := NewCommandBar(command)
	want := " j:down\t k:up\t q:quit\t"

	got := commandbar.View()

	if got != want {
		t.Errorf("Wanted \"%s\", got \"%s\"", want, got)
	}
}

func Test_Empty_Command_View(t *testing.T) {
	command := CommandMap{}
	commandbar := NewCommandBar(command)
	want := ""

	got := commandbar.View()

	if got != want {
		t.Errorf("Wanted \"%s\", got \"%s\"", want, got)
	}
}

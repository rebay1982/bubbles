package commandbar

import (
	"testing"
)

func Test_Single_Command_View(t *testing.T) {
	commands := []Command{
		Command{Key: "q", Description: "quit"},
	}
	commandbar := NewCommandBar(commands)
	want := " q:quit"

	got := commandbar.View()

	if got != want {
		t.Errorf("Wanted \"%s\", got \"%s\"", want, got)
	}
}

func Test_Multi_Command_View(t *testing.T) {
	commands := []Command{
		Command{Key: "q", Description: "quit"},
		Command{Key: "j", Description: "down"},
		Command{Key: "k", Description: "up"},
	}
	commandbar := NewCommandBar(commands)
	want := " q:quit j:down k:up"

	got := commandbar.View()

	if got != want {
		t.Errorf("Wanted \"%s\", got \"%s\"", want, got)
	}
}

func Test_Empty_Command_View(t *testing.T) {
	commands := []Command{}
	commandbar := NewCommandBar(commands)
	want := ""

	got := commandbar.View()

	if got != want {
		t.Errorf("Wanted \"%s\", got \"%s\"", want, got)
	}
}

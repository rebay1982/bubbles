package commandbar

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type Command struct {
	Key   string
	Description string
}

type CommandBar struct {
	commands []Command
}

func NewCommandBar(commands []Command) tea.Model {
	return CommandBar{
		commands: commands,
	}
}

func (c CommandBar) Init() tea.Cmd {
	return nil
}

func (c CommandBar) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	return c, nil
}

func (c CommandBar) View() string {
	var out string
	for _, cmd := range c.commands {
		out += fmt.Sprintf(" %s:%s", cmd.Key, cmd.Description)
	}

	return out
}

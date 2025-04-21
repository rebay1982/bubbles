package commandbar

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type CommandMap map[string]string

type CommandBar struct {
	commands CommandMap
}

func NewCommandBar(commands CommandMap) tea.Model {
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
	for k, v := range c.commands {
		out += fmt.Sprintf(" %s:%s\t", k, v)
	}

	return out
}

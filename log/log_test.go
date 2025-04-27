package log

import (
	"fmt"
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/rebay1982/bubbles/ansi"
)

var (
	kKey tea.Msg = tea.KeyMsg{
		Type:  tea.KeyRunes,
		Runes: []rune{'k'},
	}
	upKey tea.Msg = tea.KeyMsg{
		Type:  tea.KeyUp,
		Runes: []rune{},
	}
	jKey tea.Msg = tea.KeyMsg{
		Type:  tea.KeyRunes,
		Runes: []rune{'j'},
	}
	downKey tea.Msg = tea.KeyMsg{
		Type:  tea.KeyDown,
		Runes: []rune{},
	}
	tKey tea.Msg = tea.KeyMsg{
		Type:  tea.KeyRunes,
		Runes: []rune{'t'},
	}
)

func Test_Init(t *testing.T) {
	log, _ := NewLog("Test", 1, 1)

	got := log.Init()

	if got != nil {
		t.Errorf("Wanted nil, got \"%v\"", got)
	}
}

func Test_empty_view(t *testing.T) {
	l, _ := NewLog("Test", 1, 1)
	want := fmt.Sprintf("---[ %s ]%s\n", ansi.BoldGreen(l.title), strings.Repeat("-", l.displayWidth-7-len(l.title)))
	want += "\n"
	want += fmt.Sprintf("%s[ 100%% ]---\n", strings.Repeat("-", l.displayWidth-11))

	got := l.View()

	if got != want {
		t.Errorf("Wanted \n\"%s\", got \n\"%s\"", want, got)
	}
}

func Test_single_log_view(t *testing.T) {
	l, _ := NewLog("Test", 1, 1)
	l.Push("Log Line")
	want := fmt.Sprintf("---[ %s ]%s\n", ansi.BoldGreen(l.title), strings.Repeat("-", l.displayWidth-7-len(l.title)))
	want += fmt.Sprintf("%s %s\n", ansi.BoldWhite(">"), ansi.BoldWhite("Log Line"))
	want += fmt.Sprintf("%s[ 100%% ]---\n", strings.Repeat("-", l.displayWidth-11))

	got := l.View()

	if got != want {
		t.Errorf("Wanted \n\"%s\", got \n\"%s\"", want, got)
	}
}

func Test_multi_log_view(t *testing.T) {
	l, _ := NewLog("Test", 2, 2)
	l.Push("Log Line 1")
	l.Push("Log Line 2")
	want := fmt.Sprintf("---[ %s ]%s\n", ansi.BoldGreen(l.title), strings.Repeat("-", l.displayWidth-7-len(l.title)))
	want += fmt.Sprintf("  %s\n", "Log Line 1")
	want += fmt.Sprintf("%s %s\n", ansi.BoldWhite(">"), ansi.BoldWhite("Log Line 2"))
	want += fmt.Sprintf("%s[ 100%% ]---\n", strings.Repeat("-", l.displayWidth-11))

	got := l.View()

	if got != want {
		t.Errorf("Wanted \n\"%s\", got \n\"%s\"", want, got)
	}
}

func Test_more_display_lines_than_in_buffer_view(t *testing.T) {
	l, _ := NewLog("Test", 3, 3)
	l.Push("Log Line 1")
	want := fmt.Sprintf("---[ %s ]%s\n", ansi.BoldGreen(l.title), strings.Repeat("-", l.displayWidth-7-len(l.title)))
	want += fmt.Sprintf("%s %s\n\n\n", ansi.BoldWhite(">"), ansi.BoldWhite("Log Line 1"))
	want += fmt.Sprintf("%s[ 100%% ]---\n", strings.Repeat("-", l.displayWidth-11))

	got := l.View()

	if got != want {
		t.Errorf("Wanted \n\"%s\", got \n\"%s\"", want, got)
	}
}

func Test_push_more_lines_than_display_view_size_view(t *testing.T) {
	l, _ := NewLog("Test", 2, 3)
	l.Push("Log Line 1")
	l.Push("Log Line 2")
	l.Push("Log Line 3")
	want := fmt.Sprintf("---[ %s ]%s\n", ansi.BoldGreen(l.title), strings.Repeat("-", l.displayWidth-7-len(l.title)))
	want += fmt.Sprintf("  %s\n", "Log Line 2")
	want += fmt.Sprintf("%s %s\n", ansi.BoldWhite(">"), ansi.BoldWhite("Log Line 3"))
	want += fmt.Sprintf("%s[ 100%% ]---\n", strings.Repeat("-", l.displayWidth-11))

	got := l.View()

	if got != want {
		t.Errorf("Wanted \n\"%s\", got \n\"%s\"", want, got)
	}
}

func Test_Push_more_lines_than_buffer_size_view(t *testing.T) {
	l, _ := NewLog("Test", 2, 2)
	want := fmt.Sprintf("---[ %s ]%s\n", ansi.BoldGreen(l.title), strings.Repeat("-", l.displayWidth-7-len(l.title)))
	want += fmt.Sprintf("  %s\n", "Log Line 2")
	want += fmt.Sprintf("%s %s\n", ansi.BoldWhite(">"), ansi.BoldWhite("Log Line 3"))
	want += fmt.Sprintf("%s[ 100%% ]---\n", strings.Repeat("-", l.displayWidth-11))
	l.Push("Log Line 1")
	l.Push("Log Line 2")
	l.Push("Log Line 3")

	got := l.View()

	if got != want {
		t.Errorf("Wanted \n\"%s\", got \n\"%s\"", want, got)
	}
}

func Test_keymsg_k_update(t *testing.T) {
	nl, _ := NewLog("Test", 2, 2)
	nl.Push("Log Line 1")
	nl.Push("Log Line 2")
	l, _ := nl.Update(kKey)
	want := fmt.Sprintf("---[ %s ]%s\n", ansi.BoldGreen(nl.title), strings.Repeat("-", nl.displayWidth-7-len(nl.title)))
	want += fmt.Sprintf("%s %s\n", ansi.BoldWhite(">"), ansi.BoldWhite("Log Line 1"))
	want += fmt.Sprintf("  %s\n", "Log Line 2")
	want += fmt.Sprintf("%s[ 100%% ]---\n", strings.Repeat("-", nl.displayWidth-11))

	got := l.View()

	if got != want {
		t.Errorf("Wanted \n\"%s\", got \n\"%s\"", want, got)
	}
}

func Test_keymsg_up_update(t *testing.T) {
	nl, _ := NewLog("Test", 2, 2)
	nl.Push("Log Line 1")
	nl.Push("Log Line 2")
	l, _ := nl.Update(upKey)
	want := fmt.Sprintf("---[ %s ]%s\n", ansi.BoldGreen(nl.title), strings.Repeat("-", nl.displayWidth-7-len(nl.title)))
	want += fmt.Sprintf("%s %s\n", ansi.BoldWhite(">"), ansi.BoldWhite("Log Line 1"))
	want += fmt.Sprintf("  %s\n", "Log Line 2")
	want += fmt.Sprintf("%s[ 100%% ]---\n", strings.Repeat("-", nl.displayWidth-11))

	got := l.View()

	if got != want {
		t.Errorf("Wanted \n\"%s\", got \n\"%s\"", want, got)
	}
}

func Test_keymsg_k_j_update(t *testing.T) {
	nl, _ := NewLog("Test", 2, 2)
	nl.Push("Log Line 1")
	nl.Push("Log Line 2")
	l, _ := nl.Update(kKey)
	l, _ = l.Update(jKey)
	want := fmt.Sprintf("---[ %s ]%s\n", ansi.BoldGreen(nl.title), strings.Repeat("-", nl.displayWidth-7-len(nl.title)))
	want += fmt.Sprintf("  %s\n", "Log Line 1")
	want += fmt.Sprintf("%s %s\n", ansi.BoldWhite(">"), ansi.BoldWhite("Log Line 2"))
	want += fmt.Sprintf("%s[ 100%% ]---\n", strings.Repeat("-", nl.displayWidth-11))

	got := l.View()

	if got != want {
		t.Errorf("Wanted \n\"%s\", got \n\"%s\"", want, got)
	}
}

func Test_keymsg_k_down_update(t *testing.T) {
	nl, _ := NewLog("Test", 2, 2)
	nl.Push("Log Line 1")
	nl.Push("Log Line 2")
	l, _ := nl.Update(kKey)
	l, _ = l.Update(downKey)
	want := fmt.Sprintf("---[ %s ]%s\n", ansi.BoldGreen(nl.title), strings.Repeat("-", nl.displayWidth-7-len(nl.title)))
	want += fmt.Sprintf("  %s\n", "Log Line 1")
	want += fmt.Sprintf("%s %s\n", ansi.BoldWhite(">"), ansi.BoldWhite("Log Line 2"))
	want += fmt.Sprintf("%s[ 100%% ]---\n", strings.Repeat("-", nl.displayWidth-11))

	got := l.View()

	if got != want {
		t.Errorf("Wanted \n\"%s\", got \n\"%s\"", want, got)
	}
}
func Test_keymsg_k_k_t_update(t *testing.T) {
	nl, _ := NewLog("Test", 3, 3)
	nl.Push("Log Line 1")
	nl.Push("Log Line 2")
	nl.Push("Log Line 3")
	l, _ := nl.Update(kKey)
	l, _ = l.Update(kKey)
	l, _ = l.Update(tKey)
	want := fmt.Sprintf("---[ %s ]%s\n", ansi.BoldGreen(nl.title), strings.Repeat("-", nl.displayWidth-7-len(nl.title)))
	want += fmt.Sprintf("  %s\n", "Log Line 1")
	want += fmt.Sprintf("  %s\n", "Log Line 2")
	want += fmt.Sprintf("%s %s\n", ansi.BoldWhite(">"), ansi.BoldWhite("Log Line 3"))
	want += fmt.Sprintf("%s[ 100%% ]---\n", strings.Repeat("-", nl.displayWidth-11))

	got := l.View()

	if got != want {
		t.Errorf("Wanted \n\"%s\", got \n\"%s\"", want, got)
	}
}

func Test_keymsg_less_lines_than_buffer_size_go_to_top_update(t *testing.T) {
	nl, _ := NewLog("Test", 3, 3)
	nl.Push("Log Line 1")
	nl.Push("Log Line 2")
	l, _ := nl.Update(kKey)
	l, _ = l.Update(kKey)
	want := fmt.Sprintf("---[ %s ]%s\n", ansi.BoldGreen(nl.title), strings.Repeat("-", nl.displayWidth-7-len(nl.title)))
	want += fmt.Sprintf("%s %s\n", ansi.BoldWhite(">"), ansi.BoldWhite("Log Line 1"))
	want += fmt.Sprintf("  %s\n\n", "Log Line 2")
	want += fmt.Sprintf("%s[ 100%% ]---\n", strings.Repeat("-", nl.displayWidth-11))

	got := l.View()

	if got != want {
		t.Errorf("Wanted \n\"%s\", got \n\"%s\"", want, got)
	}
}

func Test_more_lines_than_buffer_size_go_to_top_update(t *testing.T) {
	nl, _ := NewLog("Test", 3, 3)
	nl.Push("Log Line 1")
	nl.Push("Log Line 2")
	nl.Push("Log Line 3")
	nl.Push("Log Line 4")
	nl.Push("Log Line 5")
	l, _ := nl.Update(kKey)
	l, _ = l.Update(kKey)
	l, _ = l.Update(kKey)
	l, _ = l.Update(kKey)
	want := fmt.Sprintf("---[ %s ]%s\n", ansi.BoldGreen(nl.title), strings.Repeat("-", nl.displayWidth-7-len(nl.title)))
	want += fmt.Sprintf("%s %s\n", ansi.BoldWhite(">"), ansi.BoldWhite("Log Line 3"))
	want += fmt.Sprintf("  %s\n", "Log Line 4")
	want += fmt.Sprintf("  %s\n", "Log Line 5")
	want += fmt.Sprintf("%s[ 100%% ]---\n", strings.Repeat("-", nl.displayWidth-11))

	got := l.View()

	if got != want {
		t.Errorf("Wanted \n\"%s\", got \n\"%s\"", want, got)
	}
}

func Test_view_percentage_more_lines_than_view_size_go_to_top_update(t *testing.T) {
	nl, _ := NewLog("Test", 2, 5)
	nl.Push("Log Line 1")
	nl.Push("Log Line 2")
	nl.Push("Log Line 3")
	nl.Push("Log Line 4")
	nl.Push("Log Line 5")
	l, _ := nl.Update(kKey)
	l, _ = l.Update(kKey)
	l, _ = l.Update(kKey)
	l, _ = l.Update(kKey)
	l, _ = l.Update(kKey)
	want := fmt.Sprintf("---[ %s ]%s\n", ansi.BoldGreen(nl.title), strings.Repeat("-", nl.displayWidth-7-len(nl.title)))
	want += fmt.Sprintf("%s %s\n", ansi.BoldWhite(">"), ansi.BoldWhite("Log Line 1"))
	want += fmt.Sprintf("  %s\n", "Log Line 2")
	want += fmt.Sprintf("%s[ 0%% ]---\n", strings.Repeat("-", nl.displayWidth-9))

	got := l.View()

	if got != want {
		t.Errorf("Wanted \n\"%s\", got \n\"%s\"", want, got)
	}
}

func Test_view_percentage_more_lines_than_view_size_go_up_one_stay_in_bottom_zone_update(t *testing.T) {
	nl, _ := NewLog("Test", 2, 5)
	nl.Push("Log Line 1")
	nl.Push("Log Line 2")
	nl.Push("Log Line 3")
	l, _ := nl.Update(kKey)
	want := fmt.Sprintf("---[ %s ]%s\n", ansi.BoldGreen(nl.title), strings.Repeat("-", nl.displayWidth-7-len(nl.title)))
	want += fmt.Sprintf("%s %s\n", ansi.BoldWhite(">"), ansi.BoldWhite("Log Line 2"))
	want += fmt.Sprintf("  %s\n", "Log Line 3")
	want += fmt.Sprintf("%s[ 100%% ]---\n", strings.Repeat("-", nl.displayWidth-11))

	got := l.View()

	if got != want {
		t.Errorf("Wanted \n\"%s\", got \n\"%s\"", want, got)
	}
}

func Test_view_percentage_more_lines_than_view_size_go_up_two_thirds_update(t *testing.T) {
	nl, _ := NewLog("Test", 2, 5)
	nl.Push("Log Line 1")
	nl.Push("Log Line 2")
	nl.Push("Log Line 3")
	nl.Push("Log Line 4")
	nl.Push("Log Line 5")
	l, _ := nl.Update(kKey)
	l, _ = l.Update(kKey)
	l, _ = l.Update(kKey)
	want := fmt.Sprintf("---[ %s ]%s\n", ansi.BoldGreen(nl.title), strings.Repeat("-", nl.displayWidth-7-len(nl.title)))
	want += fmt.Sprintf("%s %s\n", ansi.BoldWhite(">"), ansi.BoldWhite("Log Line 2"))
	want += fmt.Sprintf("  %s\n", "Log Line 3")
	want += fmt.Sprintf("%s[ 33%% ]---\n", strings.Repeat("-", nl.displayWidth-10))

	got := l.View()

	if got != want {
		t.Errorf("Wanted \n\"%s\", got \n\"%s\"", want, got)
	}
}

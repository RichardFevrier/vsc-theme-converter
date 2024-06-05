package wezterm

import (
	"fmt"

	"github.com/RichardFevrier/vsc-theme-converter/vscode"
)

func Generate(theme vscode.VSCodeTheme) {

	fmt.Println("\nWezterm:")

	s := "\n"
	s += "[colors]\n"
	s += "background = \"" + theme.Colors.EditorBackground + "\"\n"
	s += "foreground = \"" + theme.Colors.EditorForeground + "\"\n"
	s += "cursor_bg = \"" + theme.Colors.EditorCursorForeground + "\"\n"
	s += "cursor_border = \"" + theme.Colors.EditorCursorForeground + "\"\n"
	s += "scrollbar_thumb = \"" + theme.Colors.ScrollbarSliderBackground + "\"\n"
	s += "selection_bg = \"" + theme.Colors.EditorSelectionBackground + "\"\n"

	fmt.Println(s)
}

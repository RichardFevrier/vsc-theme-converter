package utils

import (
	"slices"
	"strings"

	"github.com/RichardFevrier/vsc-theme-converter/vscode"
)

type Token struct {
	Color, Style string
}

func Extract(theme vscode.VSCodeTheme, token string) *Token {
	for _, t := range theme.TokenColors {
		s := strings.Split(t.Scope, ", ")
		if slices.Contains(s, token) {
			return &Token{t.Settings.Foreground, t.Settings.FontStyle}
		}
	}
	return nil
}

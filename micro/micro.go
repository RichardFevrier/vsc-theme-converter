package micro

import (
	"fmt"

	"github.com/RichardFevrier/vsc-theme-converter/utils"
	"github.com/RichardFevrier/vsc-theme-converter/vscode"
)

func Generate(theme vscode.VSCodeTheme) {

	fmt.Println("\nMicro:")

	s := "\n"
	s += "color-link default \"" + theme.Colors.EditorForeground + "," + theme.Colors.EditorBackground + "\""
	append(&s, theme, "comment", "comment")
	s += "\n"
	append(&s, theme, "constant.language", "constant.bool")
	append(&s, theme, "constant.numeric", "constant.number")
	append(&s, theme, "constant.character", "constant.specialChar")
	append(&s, theme, "string", "constant.string")
	s += "\n\n"
	s += "color-link symbol.brackets \"" + "#179FFF" + "\""
	append(&s, theme, "keyword.operator", "symbol.operator")
	s += "\n"
	append(&s, theme, "storage.type", "type")
	s += "\n"
	append(&s, theme, "keyword.control", "special")
	append(&s, theme, "keyword.control", "statement")
	append(&s, theme, "keyword.other.namespace", "preproc")
	s += "\n\n"
	s += "color-link underlined \"" + theme.Colors.EditorSelectionBackground + "\"\n"
	s += "\n"
	s += "color-link diff-added \"" + theme.Colors.EditorGutterAddedBackground + "\"\n"
	s += "color-link diff-modified \"" + theme.Colors.EditorGutterModifiedBackground + "\"\n"
	s += "color-link diff-deleted \"" + theme.Colors.EditorGutterDeletedBackground + "\"\n"
	s += "\n"
	s += "color-link gutter-warning \"" + theme.Colors.EditorWarningBorder + "\"\n"
	s += "color-link gutter-error \"" + theme.Colors.EditorErrorBorder + "\""
	s += "\n"
	append(&s, theme, "token.info-token", "message")
	append(&s, theme, "token.error-token", "error-message")
	s += "\n\n"
	s += "color-link line-number \"" + theme.Colors.EditorLineNumberForeground + "\"\n"
	s += "color-link indent-char \"" + theme.Colors.EditorIndentGuideBackground + "\"\n"
	s += "color-link scrollbar \"" + theme.Colors.ScrollbarSliderBackground + "\"\n"
	s += "color-link statusline \"" + theme.Colors.ActivityBarForeground + "\"\n"
	s += "\n"
	s += "color-link cursor-line \"" + theme.Colors.EditorLineHighlightBackground + "\"\n"
	s += "color-link selection \"," + theme.Colors.EditorSelectionBackground + "\""

	fmt.Println(s)
}

func append(s *string, theme vscode.VSCodeTheme, vscToken string, microToken string) {
	t := utils.Extract(theme, vscToken)
	if t != nil {
		*s += "\ncolor-link " + microToken + " \""
		if len(t.Style) != 0 {
			*s += t.Style + " "
		}
		*s += t.Color + "\""
	}
}

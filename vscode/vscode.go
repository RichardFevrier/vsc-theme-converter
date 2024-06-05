package vscode

type VSCodeTheme struct {
	Colors struct {
		EditorBackground               string `json:"editor.background"`
		EditorForeground               string `json:"editor.foreground"`
		ScrollbarSliderBackground      string `json:"scrollbarSlider.background"`
		ActivityBarForeground          string `json:"activityBar.foreground"`
		EditorCursorForeground         string `json:"editorCursor.foreground"`
		EditorLineHighlightBackground  string `json:"editor.lineHighlightBackground"`
		EditorSelectionBackground      string `json:"editor.selectionBackground"`
		EditorLineNumberForeground     string `json:"editorLineNumber.foreground"`
		EditorIndentGuideBackground    string `json:"editorIndentGuide.background"`
		EditorGutterAddedBackground    string `json:"editorGutter.addedBackground"`
		EditorGutterDeletedBackground  string `json:"editorGutter.deletedBackground"`
		EditorGutterModifiedBackground string `json:"editorGutter.modifiedBackground"`
		EditorWarningBorder            string `json:"editorWarning.border"`
		EditorErrorBorder              string `json:"editorError.border"`
	} `json:"colors"`
	TokenColors []struct {
		Scope    string `json:"scope"`
		Settings struct {
			Foreground string `json:"foreground"`
			FontStyle  string `json:"fontStyle"`
		} `json:"settings"`
	} `json:"tokenColors"`
}

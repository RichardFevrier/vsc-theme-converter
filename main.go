package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/RichardFevrier/vsc-theme-converter/micro"
	"github.com/RichardFevrier/vsc-theme-converter/vscode"
	"github.com/RichardFevrier/vsc-theme-converter/wezterm"

	"github.com/tailscale/hujson"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: vsc-theme-converter <path_to_vscode_theme_jsonc> [output]")
		fmt.Println("[output] can be either 'all' or 'micro' or 'wezterm'. Default {all}")
		return
	}

	path := os.Args[1]
	output := "all"
	if len(os.Args) > 2 {
		output = os.Args[2]
	}

	invalidData, err := os.ReadFile(path)
	if err != nil {
		return
	}

	data, err := hujson.Standardize(invalidData)
	if err != nil {
		return
	}

	var theme vscode.VSCodeTheme

	if err := json.Unmarshal(data, &theme); err != nil {
		log.Fatal(err)
	}

	switch output {
	case "micro":
		micro.Generate(theme)
	case "wezterm":
		wezterm.Generate(theme)
	default:
		micro.Generate(theme)
		wezterm.Generate(theme)
	}
}

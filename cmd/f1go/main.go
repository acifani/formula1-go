package main

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/acifani/formula1-go/internal/ui"
)

var (
	Version = "0.1.0"

	styles = ui.NewStyles()

	rootCmd = &cobra.Command{
		Use:                   "f1go",
		Long:                  styles.Paragraph.Render("Run without arguments for a TUI or use the sub-commands like a pro."),
		DisableFlagsInUseLine: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return ui.NewProgram().Start()
		},
	}
)

func init() {
	rootCmd.Version = Version
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

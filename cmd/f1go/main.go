package main

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/acifani/formula1-go/internal/ui"
	"github.com/acifani/formula1-go/internal/ui/program"
)

var (
	version = "1.0.0"

	styles = ui.NewStyles()

	rootCmd = &cobra.Command{
		Use:                   "f1go",
		Long:                  styles.Paragraph.Render("Run without arguments for a TUI or use the sub-commands like a pro."),
		DisableFlagsInUseLine: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := program.New(styles, program.PageResults).Run()
			return err
		},
	}
)

func init() {
	rootCmd.Version = version

	rootCmd.AddGroup(&cobra.Group{ID: "commands", Title: "Commands:"})
	rootCmd.AddCommand(
		&cobra.Command{
			Use:                   "schedule",
			Aliases:               []string{"s", "season"},
			Short:                 "Season schedule",
			GroupID:               "commands",
			DisableFlagsInUseLine: true,
			RunE: func(cmd *cobra.Command, args []string) error {
				_, err := program.New(styles, program.PageSeason).Run()
				return err
			},
		},
		&cobra.Command{
			Use:                   "wcc",
			Aliases:               []string{"c"},
			Short:                 "Constructor Standings",
			GroupID:               "commands",
			DisableFlagsInUseLine: true,
			RunE: func(cmd *cobra.Command, args []string) error {
				_, err := program.New(styles, program.PageWCC).Run()
				return err
			},
		},
		&cobra.Command{
			Use:                   "wdc",
			Aliases:               []string{"d"},
			Short:                 "Driver Standings",
			GroupID:               "commands",
			DisableFlagsInUseLine: true,
			RunE: func(cmd *cobra.Command, args []string) error {
				_, err := program.New(styles, program.PageWDC).Run()
				return err
			},
		},
	)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

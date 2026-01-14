package cmd

import (
	"github.com/jiangtao/ccconfig/pkg/config"
	"github.com/jiangtao/ccconfig/pkg/i18n"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ccconfig",
	Short: "Claude Code configuration backup/restore tool",
	Long: `ccconfig is a CLI tool for backing up and restoring Claude Code configurations.

It supports:
  • Backing up settings, commands, skills, and project configs
  • Restoring configurations to new machines
  • Managing plugin caches
  • Internationalization (en/zh)`,
}

// Execute runs the root command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Global flags
	rootCmd.PersistentFlags().StringP("lang", "l", "", "Language (en/zh)")
	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if lang, _ := cmd.Flags().GetString("lang"); lang != "" {
			config.SetLang(lang)
			if err := i18n.Init(lang); err != nil {
				return err
			}
		}
		return nil
	}
}

// GetRootCommand returns the root command for adding subcommands
func GetRootCommand() *cobra.Command {
	return rootCmd
}

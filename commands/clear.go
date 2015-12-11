package commands

import (
	"github.com/khlieng/dispatch/Godeps/_workspace/src/github.com/spf13/cobra"

	"github.com/khlieng/dispatch/storage"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all application data",
	Run: func(cmd *cobra.Command, args []string) {
		storage.Clear(appDir)
	},
}

/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// volumeCmd represents the "volume" command in the CLI.
// This command will be the parent for volume-related subcommands like snapshots.
var volumeCmd = &cobra.Command{
	Use:   "volume",
	Short: "Manage Kubernetes volumes and snapshots",
	Long: `The "volume" command allows you to perform operations on Kubernetes volumes.
You can create, snapshot, restore, or list volumes using its subcommands.

Example usage:
  k8backup volume snapshot
  k8backup volume restore
`,
	Run: func(cmd *cobra.Command, args []string) {
		// This will be executed if "volume" is called without subcommands
		log.Println("[INFO] Volume command called. Use a subcommand like 'snapshot' or 'restore'.")
		fmt.Println("volume command invoked")
	},
}

func init() {
	// Add the volume command as a child of the root command
	rootCmd.AddCommand(volumeCmd)
}

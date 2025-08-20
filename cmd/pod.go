/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// podCmd is the parent command for managing Kubernetes Pods.
// Subcommands like 'backup' and 'restore' are attached to it.
var podCmd = &cobra.Command{
	Use:   "pod",
	Short: "Manage Kubernetes Pods",
	Long: `The 'pod' command allows you to perform backup and restore operations
on Kubernetes Pod resources. Use its subcommands to backup or restore pods.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("[INFO] Pod command called. Use subcommands like 'backup' or 'restore'.")
	},
}

func init() {
	// Attach podCmd to the root command
	rootCmd.AddCommand(podCmd)
}

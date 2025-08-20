/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// Struct placeholder for deployment command flags (future use)
type deploymentCmdFlags struct {
	// future flags here
}

var depCmdFlags deploymentCmdFlags


// deploymentCmd is the parent command for managing Kubernetes Deployments.
// Subcommands like 'backup' and 'restore' are attached to it.
var deploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "Manage Kubernetes Deployments",
	Long: `The 'deployment' command allows you to perform backup and restore operations
on Kubernetes Deployment resources. Use its subcommands to backup or restore deployments.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("[INFO] Deployment command called. Use subcommands like 'backup' or 'restore'.")
	},
}

func init() {
	// Attach deploymentCmd to the root command
	rootCmd.AddCommand(deploymentCmd)
}

/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands.
// All other commands (backup, restore, volume, snapshot, etc.) are added as subcommands to rootCmd.
var rootCmd = &cobra.Command{
	Use:   "K8Backup",
	Short: "A CLI tool to backup and restore Kubernetes resources",
	Long: `K8Backup is a CLI tool that simplifies the management of Kubernetes resources.
It allows you to create backups of Deployments, Pods, PVCs, and Volumes, as well as restore them with minimal effort.
The tool abstracts the complexities of interacting with the Kubernetes API directly.`,
}

// Execute runs the root command and handles any errors.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("[ERROR] Failed to execute root command: %v", err)
		os.Exit(1)
	}
}

func init() {
	// Define global flags for the CLI
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"K8Backup/objects"
	"github.com/spf13/cobra"
)

// listCmd represents the "list" command in the CLI.
// It lists all backup objects currently stored by K8Backup.
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all backup objects",
	Long: `The 'list' command displays all backup objects created by K8Backup.
It shows information such as resource name, namespace, type, and backup file path.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("[INFO] Listing all backups...")
		objects.ListBackups()
		log.Println("[SUCCESS] Backup list displayed")
	},
}

func init() {
	// Attach listCmd to the root command
	rootCmd.AddCommand(listCmd)
}

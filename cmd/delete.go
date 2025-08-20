/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"K8Backup/objects"
	"github.com/spf13/cobra"
)

// Struct to hold flags for the delete command
type deleteFlags struct {
	FilePath string
}

var delFlags deleteFlags

// deleteCmd represents the "delete" command in the CLI.
// It deletes a backup object by specifying the backup YAML or object file.
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a backup object",
	Long: `Deletes a backup object created by K8Backup.
You must provide the file path of the backup object using the --file flag.

Example:
  k8backup delete --file backups/mypod-default-pod.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		if delFlags.FilePath == "" {
			log.Fatal("[ERROR] Please provide a valid file path using --file")
		}
		objects.DeleteBackup(delFlags.FilePath)
		log.Printf("[SUCCESS] Backup deleted: %s", delFlags.FilePath)
	},
}

func init() {
	// Attach deleteCmd to the root command
	rootCmd.AddCommand(deleteCmd)

	// Define flag for specifying backup file to delete
	deleteCmd.Flags().StringVarP(&delFlags.FilePath, "file", "f", "", "Path to the backup file to delete")
}

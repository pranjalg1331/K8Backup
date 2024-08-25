/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"K8Backup/objects"
	"github.com/spf13/cobra"
)

var backupfilepath string;

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete any backup object by passing the objects file name in --file flag",
	
	Run: func(cmd *cobra.Command, args []string) {
		objects.DeleteBackup(backupfilepath)
		fmt.Println("delete called")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringVarP(&backupfilepath, "file","", "", "Name of the resource")

}

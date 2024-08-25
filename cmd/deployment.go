/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)




// deploymentCmd represents the deployment command
var deploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "Backup and Restore Deployment",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deployment called")
	},
}

func init() {
	rootCmd.AddCommand(deploymentCmd)

}

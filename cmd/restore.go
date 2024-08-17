/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"K8Backup/internal"
	"github.com/spf13/cobra"
)

var backupObjectName string

var restoredeploymentCmd = &cobra.Command{
	Use:   "restore",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		clientset:=internal.Connect(path)
		internal.RestoreDeployment(clientset,backupObjectName)
		
		fmt.Println("restoredeployment called")
	},
}


// restorepodCmd represents the restorepod command
var restorepodCmd = &cobra.Command{
	Use:   "restore",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		clientset:=internal.Connect(path)
		internal.RestorePod(clientset,backupObjectName)
		fmt.Println("restorepod called")
	},
}

func init() {
	podCmd.AddCommand(restorepodCmd)
	deploymentCmd.AddCommand(restoredeploymentCmd)

	restorepodCmd.Flags().StringVarP(&backupObjectName, "object", "o", "", "Name of the resource")
	restorepodCmd.Flags().StringVarP(&path, "path", "p", "", "Name of the resource")  

	restoredeploymentCmd.Flags().StringVarP(&backupObjectName, "object", "o", "", "Name of the resource")
	restoredeploymentCmd.Flags().StringVarP(&path, "path", "p", "", "Name of the resource")  

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// restorepodCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// restorepodCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"K8Backup/internal"
	"github.com/spf13/cobra"
)

var backupObjectName,restorename string

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
		internal.RestoreDeployment(clientset,backupObjectName,restorename)
		
		fmt.Println("restoredeployment called")
	},
}


// restorepodCmd represents the restorepod command
var restorepodCmd = &cobra.Command{
	Use:   "restore",
	Short: "restore pod by passing file path of ",
	Run: func(cmd *cobra.Command, args []string) {
		clientset:=internal.Connect(path)
		internal.RestorePod(clientset,backupObjectName,restorename)
		fmt.Println("restorepod called")
	},
}


var restorePvcCmd = &cobra.Command{
	Use:   "restore",
	Short: "restore pvc from volume snapshot",
	Run: func(cmd *cobra.Command, args []string) {
		clientset:=internal.Connect(path)
		internal.RestorePvc(clientset,backupObjectName,restorename)
		fmt.Println("restorepvc called")
	},
}




func init() {
	podCmd.AddCommand(restorepodCmd)
	deploymentCmd.AddCommand(restoredeploymentCmd)
	volumeCmd.AddCommand(restorePvcCmd)

	restorepodCmd.Flags().StringVarP(&backupObjectName, "object", "o", "", "Name of the resource")
	restorepodCmd.Flags().StringVarP(&path, "path", "p", "", "Name of the resource")  
	restorepodCmd.Flags().StringVarP(&restorename, "name", "n", "", "Name of the resource")  


	restoredeploymentCmd.Flags().StringVarP(&backupObjectName, "object", "o", "", "Name of the resource")
	restoredeploymentCmd.Flags().StringVarP(&path, "path", "p", "", "Name of the resource")  
	restoredeploymentCmd.Flags().StringVarP(&restorename, "name", "n", "", "Name of the resource")  

	restorePvcCmd.Flags().StringVarP(&backupObjectName, "object", "o", "", "Name of the resource")
	restorePvcCmd.Flags().StringVarP(&path, "path", "p", "", "Name of the resource") 
	restorePvcCmd.Flags().StringVarP(&restorename, "name", "n", "", "Name of the resource")   
}

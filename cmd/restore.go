/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"K8Backup/internal"
	"github.com/spf13/cobra"
)

// Structs to hold flags for restore commands
type restoreFlags struct {
	ObjectName string
	Name       string
	Path       string
}

var restorePodFlags, restoreDeploymentFlags, restorePvcFlags restoreFlags

// restoredeploymentCmd restores a Deployment from a backup
var restoredeploymentCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore a deployment from backup",
	Long:  "Restores a deployment using a backup YAML file and kubeconfig path.",
	Run: func(cmd *cobra.Command, args []string) {
		clientset := internal.Connect(restoreDeploymentFlags.Path)
		internal.RestoreDeployment(clientset, restoreDeploymentFlags.ObjectName, restoreDeploymentFlags.Name)
		fmt.Println("[SUCCESS] Deployment restored")
	},
}

// restorepodCmd restores a Pod from a backup
var restorepodCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore a pod from backup",
	Run: func(cmd *cobra.Command, args []string) {
		clientset := internal.Connect(restorePodFlags.Path)
		internal.RestorePod(clientset, restorePodFlags.ObjectName, restorePodFlags.Name)
		fmt.Println("[SUCCESS] Pod restored")
	},
}

// restorePvcCmd restores a PVC from a snapshot
var restorePvcCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore a PVC from volume snapshot",
	Run: func(cmd *cobra.Command, args []string) {
		clientset := internal.Connect(restorePvcFlags.Path)
		internal.RestorePvc(clientset, restorePvcFlags.ObjectName, restorePvcFlags.Name)
		fmt.Println("[SUCCESS] PVC restored")
	},
}

func init() {
	podCmd.AddCommand(restorepodCmd)
	deploymentCmd.AddCommand(restoredeploymentCmd)
	volumeCmd.AddCommand(restorePvcCmd)

	// Flags for restorepodCmd
	restorepodCmd.Flags().StringVarP(&restorePodFlags.ObjectName, "object", "o", "", "Name of the resource")
	restorepodCmd.Flags().StringVarP(&restorePodFlags.Path, "path", "p", "", "Path to kubeconfig file")
	restorepodCmd.Flags().StringVarP(&restorePodFlags.Name, "name", "n", "", "Name for the restored resource")

	// Flags for restoredeploymentCmd
	restoredeploymentCmd.Flags().StringVarP(&restoreDeploymentFlags.ObjectName, "object", "o", "", "Name of the resource")
	restoredeploymentCmd.Flags().StringVarP(&restoreDeploymentFlags.Path, "path", "p", "", "Path to kubeconfig file")
	restoredeploymentCmd.Flags().StringVarP(&restoreDeploymentFlags.Name, "name", "n", "", "Name for the restored resource")

	// Flags for restorePvcCmd
	restorePvcCmd.Flags().StringVarP(&restorePvcFlags.ObjectName, "object", "o", "", "Name of the resource")
	restorePvcCmd.Flags().StringVarP(&restorePvcFlags.Path, "path", "p", "", "Path to kubeconfig file")
	restorePvcCmd.Flags().StringVarP(&restorePvcFlags.Name, "name", "n", "", "Name for the restored PVC")
}

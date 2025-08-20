// cmd/backup.go
package cmd

import (
	"log"

	"K8Backup/internal"
	"github.com/spf13/cobra"
)

// Structs to hold flags for each backup command
type podBackupFlags struct {
	Name      string
	Namespace string
	Path      string
}

type deploymentBackupFlags struct {
	Name      string
	Namespace string
	Path      string
}

type volumeBackupFlags struct {
	Name       string
	Namespace  string
	Path       string
	BackupName string
}

var podFlags podBackupFlags
var deploymentFlags deploymentBackupFlags
var volumeFlags volumeBackupFlags

// backupPodCmd creates a backup of a Pod resource
var backupPodCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup a Pod resource",
	Long:  "Creates a backup of a specified Pod resource in a Kubernetes cluster by exporting its YAML.",
	Run: func(cmd *cobra.Command, args []string) {
		clientset := internal.Connect(podFlags.Path)
		internal.SavePod(clientset, podFlags.Name, podFlags.Namespace)
		log.Printf("[SUCCESS] Backup created for Pod '%s' in namespace '%s'", podFlags.Name, podFlags.Namespace)
	},
}

// backupDeploymentCmd creates a backup of a Deployment resource
var backupDeploymentCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup a Deployment resource",
	Long:  "Creates a backup of a specified Deployment resource in a Kubernetes cluster by exporting its YAML.",
	Run: func(cmd *cobra.Command, args []string) {
		clientset := internal.Connect(deploymentFlags.Path)
		internal.SaveDeployment(clientset, deploymentFlags.Name, deploymentFlags.Namespace)
		log.Printf("[SUCCESS] Backup created for Deployment '%s' in namespace '%s'", deploymentFlags.Name, deploymentFlags.Namespace)
	},
}

// snapshotVolumeCmd creates a VolumeSnapshot for a PVC
var snapshotVolumeCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup a PVC by creating a VolumeSnapshot",
	Long:  "Creates a snapshot of a PersistentVolumeClaim (PVC) to enable backup and restore of volume data.",
	Run: func(cmd *cobra.Command, args []string) {
		internal.CreateVolumeSnapshot(volumeFlags.Name, volumeFlags.Namespace, volumeFlags.Path, volumeFlags.BackupName)
		log.Printf("[SUCCESS] VolumeSnapshot '%s' created for PVC '%s' in namespace '%s'", volumeFlags.BackupName, volumeFlags.Name, volumeFlags.Namespace)
	},
}

func init() {
	// Attach backup commands to their respective parent commands
	podCmd.AddCommand(backupPodCmd)
	deploymentCmd.AddCommand(backupDeploymentCmd)
	volumeCmd.AddCommand(snapshotVolumeCmd)

	// Flags for backupPodCmd
	backupPodCmd.Flags().StringVarP(&podFlags.Name, "name", "n", "", "Name of the Pod to backup")
	backupPodCmd.Flags().StringVarP(&podFlags.Namespace, "namespace", "s", "default", "Namespace of the Pod")
	backupPodCmd.Flags().StringVarP(&podFlags.Path, "path", "p", "", "Path to kubeconfig file")

	// Flags for backupDeploymentCmd
	backupDeploymentCmd.Flags().StringVarP(&deploymentFlags.Name, "name", "n", "", "Name of the Deployment to backup")
	backupDeploymentCmd.Flags().StringVarP(&deploymentFlags.Namespace, "namespace", "s", "default", "Namespace of the Deployment")
	backupDeploymentCmd.Flags().StringVarP(&deploymentFlags.Path, "path", "p", "", "Path to kubeconfig file")

	// Flags for snapshotVolumeCmd
	snapshotVolumeCmd.Flags().StringVarP(&volumeFlags.Name, "name", "n", "", "Name of the PVC to snapshot")
	snapshotVolumeCmd.Flags().StringVarP(&volumeFlags.Namespace, "namespace", "s", "default", "Namespace of the PVC")
	snapshotVolumeCmd.Flags().StringVarP(&volumeFlags.Path, "path", "p", "", "Path to kubeconfig file")
	snapshotVolumeCmd.Flags().StringVarP(&volumeFlags.BackupName, "backup", "b", "", "Name for the VolumeSnapshot")
}

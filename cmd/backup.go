// cmd/backup.go
package cmd

import (
    "fmt"
    "K8Backup/internal"
    "github.com/spf13/cobra"
)

var name, namespace, path,backupname string

// backupPodCmd represents the backup command
var backupPodCmd = &cobra.Command{
    Use:   "backup",
    Short: "Backup a pod resource",
    Long:  "Creates a backup of a specified pod resource in a Kubernetes cluster.",
    Run: func(cmd *cobra.Command, args []string) {
        clientset:=internal.Connect(path)
		internal.SavePod(clientset,name,namespace)
        fmt.Println("Backup created:")
         
    },
}

var backupDeploymentCmd = &cobra.Command{
	Use:   "backup",
	Short: "backup k8s deployment",

	Run: func(cmd *cobra.Command, args []string) {
		clientset:=internal.Connect(path)
        internal.SaveDeployment(clientset,name,namespace)
        fmt.Println("Backup created:")
	},
}

var snapshotVolumeCmd = &cobra.Command{
	Use:   "backup",
	Short: "backup k8s volume by creating volume snapshots",
	Run: func(cmd *cobra.Command, args []string) {
        internal.CreateVolumeSnapshot(name,namespace,path,backupname);
        fmt.Println("Snapshot created")
	},
}


func init() {
    podCmd.AddCommand(backupPodCmd)

    // Flags for the backup command
    backupPodCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the resource")
    backupPodCmd.Flags().StringVarP(&namespace, "namespace", "s", "default", "Namespace of the resource")
    backupPodCmd.Flags().StringVarP(&path, "path", "p", "", "Name of the resource")  


    deploymentCmd.AddCommand(backupDeploymentCmd)

	backupDeploymentCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the resource")
    backupDeploymentCmd.Flags().StringVarP(&namespace, "namespace", "s", "default", "Namespace of the resource")
    backupDeploymentCmd.Flags().StringVarP(&path, "path", "p", "", "Name of the resource")  
  

    volumeCmd.AddCommand(snapshotVolumeCmd)

    snapshotVolumeCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the resource")
    snapshotVolumeCmd.Flags().StringVarP(&namespace, "namespace", "s", "default", "Namespace of the resource")
    snapshotVolumeCmd.Flags().StringVarP(&path, "path", "p", "", "Name of the resource")  
    snapshotVolumeCmd.Flags().StringVarP(&backupname, "backup", "b", "", "Name of the resource")  



    
}

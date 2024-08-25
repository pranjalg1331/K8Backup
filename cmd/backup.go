// cmd/backup.go
package cmd

import (
    "fmt"
    "K8Backup/internal"
    "github.com/spf13/cobra"
)

var name, namespace, path string

// backupPodCmd represents the backup command
var backupPodCmd = &cobra.Command{
    Use:   "backup",
    Short: "Backup a pod resource",
    Long:  "Creates a backup of a specified pod resource in a Kubernetes cluster.",
    Run: func(cmd *cobra.Command, args []string) {
        clientset:=internal.Connect(path)
		_=internal.SavePod(clientset,name,namespace)
        fmt.Println("Backup created:")
         
    },
}

var backupDeploymentCmd = &cobra.Command{
	Use:   "backup",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		clientset:=internal.Connect(path)
        _=internal.SaveDeployment(clientset,name,namespace)
        fmt.Println("Backup created:")
	},
}

var snapshotVolumeCmd = &cobra.Command{
	Use:   "backup",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
        internal.CreateVolumeSnapshot(name,namespace,path);
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



    
}

// cmd/backup.go
package cmd

import (
    "fmt"
    "K8Backup/internal"
    "github.com/spf13/cobra"
)

var name, namespace, path string

// backupCmd represents the backup command
var backupCmd = &cobra.Command{
    Use:   "backup",
    Short: "Backup a pod resource",
    Long:  "Creates a backup of a specified pod resource in a Kubernetes cluster.",
    Run: func(cmd *cobra.Command, args []string) {
        clientset:=internal.Connect(path)
		_=internal.SavePod(clientset,name,namespace)
        fmt.Println("Backup created:")
         
    },
}

func init() {
    podCmd.AddCommand(backupCmd)

    // Flags for the backup command
    backupCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the resource")
    backupCmd.Flags().StringVarP(&namespace, "namespace", "s", "default", "Namespace of the resource")
    backupCmd.Flags().StringVarP(&path, "path", "p", "", "Name of the resource")  

    
}

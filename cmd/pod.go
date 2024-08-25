// cmd/pod.go
package cmd

import (
    "github.com/spf13/cobra"
    // "cmd/rootCmd"
)

// podCmd represents the pod command
var podCmd = &cobra.Command{
    Use:   "pod",
    Short: "Backup and Restore Pod",
}

func init() {
    rootCmd.AddCommand(podCmd)
}

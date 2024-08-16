// cmd/pod.go
package cmd

import (
    "github.com/spf13/cobra"
    // "cmd/rootCmd"
)

// podCmd represents the pod command
var podCmd = &cobra.Command{
    Use:   "pod",
    Short: "Manage pod resources",
    Long:  "Commands for managing pod resources in a Kubernetes cluster.",
}

func init() {
    rootCmd.AddCommand(podCmd)
}

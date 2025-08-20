package internal

import (
	"log"

	kubernetes "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	clientcmd "k8s.io/client-go/tools/clientcmd"
)

// Connect initializes and returns a Kubernetes clientset.
// It takes a path to the kubeconfig file as input.
func Connect(path string) *kubernetes.Clientset {
	log.Printf("[INFO] Using kubeconfig from: %s", path)

	// Get Kubernetes REST config from kubeconfig file
	config := GetConfig(path)

	// Create clientset (main entry point for K8s API calls)
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Panicf("[ERROR] Failed to create K8s clientset: %v", err)
	}

	log.Println("[INFO] Successfully created Kubernetes clientset")
	return clientset
}

// GetConfig loads Kubernetes configuration from a kubeconfig file.
// Returns a REST config that can be used to create a clientset.
func GetConfig(path string) *rest.Config {
	// BuildConfigFromFlags("", path) loads config from the specified kubeconfig file.
	// If path is empty, it usually defaults to ~/.kube/config
	config, err := clientcmd.BuildConfigFromFlags("", path)
	if err != nil {
		log.Panicf("[ERROR] Failed to load K8s config from %s: %v", path, err)
	}

	log.Printf("[INFO] Successfully loaded K8s config from: %s", path)
	return config
}

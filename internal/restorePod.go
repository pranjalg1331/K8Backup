package internal

import (
	"context"
	"log"
	"os"
	"path/filepath"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"

	"gopkg.in/yaml.v3"
)

// RestorePod restores a Pod from a YAML manifest on disk.
// - clientset: Kubernetes clientset
// - fileName: path to the YAML file containing the original Pod spec
// - restorename: new name to assign to the restored Pod
func RestorePod(clientset *kubernetes.Clientset, fileName, restorename string) {
	log.Printf("[INFO] Starting restore process for Pod from file: %s", fileName)

	// Read YAML file
	filePath := filepath.Join(fileName)
	yamlData, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("[ERROR] Failed to read YAML file %s: %v", filePath, err)
	}

	// Unmarshal YAML into Pod struct
	var newpod v1.Pod
	err = yaml.Unmarshal(yamlData, &newpod)
	if err != nil {
		log.Fatalf("[ERROR] Failed to unmarshal YAML into Pod: %v", err)
	}

	// Define new Pod object with a new name but same spec
	newPod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      restorename,    // use provided restore name
			Namespace: newpod.Namespace,
			Labels:    newpod.Labels,
		},
		Spec: newpod.Spec,
	}

	// Create the Pod in the given namespace
	_, err = clientset.CoreV1().Pods(newpod.Namespace).Create(context.TODO(), newPod, metav1.CreateOptions{})
	if err != nil {
		log.Fatalf("[ERROR] Failed to create restored Pod %s in namespace %s: %v", restorename, newpod.Namespace, err)
	}

	log.Printf("[SUCCESS] Pod %s restored successfully in namespace %s", restorename, newpod.Namespace)
}

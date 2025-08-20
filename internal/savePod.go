package internal

import (
	"fmt"
	"log"
	"os"
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
	"gopkg.in/yaml.v3"
	"path/filepath"
	"K8Backup/objects"
)

// SavePod fetches a Pod from the cluster and saves its spec to a YAML file.
// - clientset: Kubernetes clientset
// - podName: name of the Pod to back up
// - namespace: namespace of the Pod
func SavePod(clientset *kubernetes.Clientset, podName string, namespace string) {
	log.Printf("[INFO] Starting backup for Pod '%s' in namespace '%s'...", podName, namespace)

	// Get Pod from Kubernetes API
	podsClient := clientset.CoreV1().Pods(namespace)
	originalPod, err := podsClient.Get(context.TODO(), podName, metav1.GetOptions{})
	if err != nil {
		log.Fatalf("[ERROR] Failed to fetch Pod '%s': %v", podName, err)
	}

	// Marshal Pod into YAML
	podYAML, err := yaml.Marshal(originalPod)
	if err != nil {
		log.Fatalf("[ERROR] Failed to marshal Pod '%s' into YAML: %v", podName, err)
	}

	// Build backup file path (backups/<name>-<namespace>-pod.yaml)
	fileName := fmt.Sprintf("%s-%s-pod.yaml", podName, namespace)
	filePath := filepath.Join("backups", fileName)

	// Save YAML to file
	err = os.WriteFile(filePath, podYAML, 0644)
	if err != nil {
		log.Fatalf("[ERROR] Failed to save Pod YAML for '%s': %v", podName, err)
	}

	// Register backup in objects package
	backup, _ := objects.CreateBackup(podName, namespace, "pod", filePath)
	if backup != nil {
		log.Printf("[SUCCESS] Pod '%s' backup saved at %s", podName, filePath)
	}
}

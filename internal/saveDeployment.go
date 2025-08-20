package internal

import (
	"fmt"
	"log"
	"os"
	"context"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
	"gopkg.in/yaml.v3"
	"path/filepath"
	"K8Backup/objects"
)

// SaveDeployment fetches a Deployment from the cluster and saves its spec to a YAML file.
// - clientset: Kubernetes clientset
// - depName: name of the Deployment to back up
// - namespace: namespace of the Deployment
func SaveDeployment(clientset *kubernetes.Clientset, depName string, namespace string) {
	log.Printf("[INFO] Starting backup for Deployment '%s' in namespace '%s'...", depName, namespace)

	// Get Deployment from Kubernetes API
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault) // NOTE: currently hardcoded to "default"
	dep, err := deploymentsClient.Get(context.TODO(), depName, metav1.GetOptions{})
	if err != nil {
		log.Fatalf("[ERROR] Failed to fetch Deployment '%s': %v", depName, err)
	}

	// Marshal Deployment into YAML
	depYAML, err := yaml.Marshal(dep)
	if err != nil {
		log.Fatalf("[ERROR] Failed to marshal Deployment '%s' into YAML: %v", depName, err)
	}

	// Build backup file path (backups/<name>-<namespace>-dep.yaml)
	fileName := fmt.Sprintf("%s-%s-dep.yaml", depName, namespace)
	filePath := filepath.Join("backups", fileName)

	// Save YAML to file
	err = os.WriteFile(filePath, depYAML, 0644)
	if err != nil {
		log.Fatalf("[ERROR] Failed to save Deployment YAML for '%s': %v", depName, err)
	}

	// Register backup in objects package
	backup, _ := objects.CreateBackup(depName, namespace, "deployment", filePath)
	if backup != nil {
		log.Printf("[SUCCESS] Deployment '%s' backup saved at %s", depName, filePath)
	}
}

package internal

import (
	"context"
	"log"
	"os"
	"path/filepath"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"

	"gopkg.in/yaml.v3"
)

// RestoreDeployment restores a Deployment from a YAML manifest on disk.
// - clientset: Kubernetes clientset
// - fileName: path to the YAML file containing the original Deployment spec
// - restorename: name to give to the restored Deployment
func RestoreDeployment(clientset *kubernetes.Clientset, fileName, restorename string) {
	log.Printf("[INFO] Starting restore process from file: %s", fileName)

	// Kubernetes client for Deployments (default namespace)
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

	// Read YAML file
	filePath := filepath.Join(fileName)
	yamlData, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("[ERROR] Failed to read YAML file %s: %v", filePath, err)
	}

	// Unmarshal YAML into Deployment struct
	var newDep appsv1.Deployment
	err = yaml.Unmarshal(yamlData, &newDep)
	if err != nil {
		log.Fatalf("[ERROR] Failed to unmarshal YAML into Deployment: %v", err)
	}

	// Construct new Deployment object with the provided restore name
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: restorename, // use new name instead of overwriting existing one
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: newDep.Spec.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: newDep.Spec.Selector.MatchLabels,
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: newDep.Spec.Template.Labels,
				},
				Spec: newDep.Spec.Template.Spec,
			},
		},
	}

	// Create the restored Deployment in the cluster
	_, err = deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		log.Fatalf("[ERROR] Failed to create restored Deployment %s: %v", restorename, err)
	}

	log.Printf("[SUCCESS] Deployment %s restored successfully from file %s", restorename, fileName)
}

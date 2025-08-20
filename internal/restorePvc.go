package internal

import (
	"context"
	"log"

	"k8s.io/apimachinery/pkg/api/resource"
	k8api "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
)

// RestorePvc creates a new PVC from a given VolumeSnapshot.
// - clientset: Kubernetes clientset
// - objectName: name of the existing VolumeSnapshot to restore from
// - restorename: name for the new PVC to be created
func RestorePvc(clientset *kubernetes.Clientset, objectName string, restorename string) {
	log.Printf("[INFO] Starting restore of PVC '%s' from snapshot '%s'...", restorename, objectName)

	// Define new PVC spec pointing to a snapshot as the data source
	pvc := &k8api.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name:      restorename,
			Namespace: "default", // TODO: make namespace configurable
		},
		Spec: k8api.PersistentVolumeClaimSpec{
			AccessModes: []k8api.PersistentVolumeAccessMode{
				k8api.ReadWriteOnce, // commonly used access mode
			},
			Resources: k8api.VolumeResourceRequirements{
				Requests: k8api.ResourceList{
					k8api.ResourceStorage: resource.MustParse("10Gi"), // requested storage size
				},
			},
			DataSource: &k8api.TypedLocalObjectReference{
				APIGroup: &[]string{"snapshot.storage.k8s.io"}[0], // API group for CSI snapshots
				Kind:     "VolumeSnapshot",                        // referencing a snapshot
				Name:     objectName,                              // snapshot to restore from
			},
			// StorageClassName: &[]string{"my-storage-class"}[0], // optional: use specific StorageClass
		},
	}

	// Create the PVC
	pvcClient := clientset.CoreV1().PersistentVolumeClaims("default")
	_, err := pvcClient.Create(context.TODO(), pvc, metav1.CreateOptions{})
	if err != nil {
		log.Fatalf("[ERROR] Failed to create PVC %s: %v", restorename, err)
	}

	log.Printf("[SUCCESS] PVC '%s' created successfully from snapshot '%s'", restorename, objectName)
}

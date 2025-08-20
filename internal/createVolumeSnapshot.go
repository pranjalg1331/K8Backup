package internal

import (
	"context"
	"log"

	apiv1 "github.com/kubernetes-csi/external-snapshotter/client/v4/apis/volumesnapshot/v1"
	v1 "github.com/kubernetes-csi/external-snapshotter/client/v4/clientset/versioned/typed/volumesnapshot/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreateVolumeSnapshot creates a VolumeSnapshot for a given PVC.
// - pvcname: name of the PersistentVolumeClaim to snapshot
// - namespace: namespace where the PVC exists
// - path: path to the kubeconfig file
// - backupname: name to give the snapshot
func CreateVolumeSnapshot(pvcname string, namespace string, path string, backupname string) {
	log.Printf("[INFO] Starting snapshot creation for PVC: %s in namespace: %s", pvcname, namespace)

	// Load Kubernetes config from kubeconfig path
	config := GetConfig(path)

	// SnapshotClass to use (must exist in the cluster)
	snapshotClassName := "csi-hostpath-snapclass"

	// Define the VolumeSnapshot object
	volumeSnapshot := &apiv1.VolumeSnapshot{
		ObjectMeta: metav1.ObjectMeta{
			Name:      backupname,
			Namespace: namespace,
		},
		Spec: apiv1.VolumeSnapshotSpec{
			VolumeSnapshotClassName: &snapshotClassName,
			Source: apiv1.VolumeSnapshotSource{
				PersistentVolumeClaimName: &pvcname,
			},
		},
	}

	// Create snapshot client
	snapshotClient, err := v1.NewForConfig(config)
	if err != nil {
		log.Fatalf("[ERROR] Failed to create snapshot client: %v", err)
	}

	// Create the snapshot in the given namespace
	_, err = snapshotClient.VolumeSnapshots(namespace).Create(context.TODO(), volumeSnapshot, metav1.CreateOptions{})
	if err != nil {
		log.Fatalf("[ERROR] Failed to create snapshot %s: %v", backupname, err)
	}

	log.Printf("[SUCCESS] Snapshot %s created for PVC %s in namespace %s", backupname, pvcname, namespace)
}

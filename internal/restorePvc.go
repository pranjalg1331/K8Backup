package internal

import (

	"context"
	"log"

	"k8s.io/apimachinery/pkg/api/resource"

	k8api "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
	
)

func RestorePvc(clientset *kubernetes.Clientset,objectName string,restorename string){

	// clientset:=Connect(path)


	pvc:=&k8api.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name:      restorename,
			Namespace: "default",
		},
		Spec: k8api.PersistentVolumeClaimSpec{
			AccessModes: []k8api.PersistentVolumeAccessMode{
				k8api.ReadWriteOnce,
			},
			Resources: k8api.VolumeResourceRequirements{
				Requests: k8api.ResourceList{
					
					k8api.ResourceStorage: resource.MustParse("10Gi"),
					// ResourceStorage:["10Gi"],
				},
			},
			DataSource: &k8api.TypedLocalObjectReference{
				APIGroup: &[]string{"snapshot.storage.k8s.io"}[0],
				Kind:     "VolumeSnapshot",
				Name:     objectName,
			},
			// StorageClassName: &[]string{"my-storage-class"}[0],
		},
		


	}

	

	pvcClient := clientset.CoreV1().PersistentVolumeClaims("default")
	_, err := pvcClient.Create(context.TODO(), pvc, metav1.CreateOptions{})
	if err != nil {
		log.Fatalf("Failed to create PVC: %v", err)
	}


	

	
}

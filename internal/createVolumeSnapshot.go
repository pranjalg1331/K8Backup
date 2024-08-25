package internal

import (
	"context"
	"log"
	v1 "github.com/kubernetes-csi/external-snapshotter/client/v4/clientset/versioned/typed/volumesnapshot/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apiv1 "github.com/kubernetes-csi/external-snapshotter/client/v4/apis/volumesnapshot/v1"

)

func CreateVolumeSnapshot(pvcname string,namespace string,path string){

	config:= GetConfig(path)
	snapshotclassname:="csi-hostpath-snapclass"

	volumeSnapshot := &apiv1.VolumeSnapshot{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "golang-snapshot-2",
			Namespace: namespace,
		},
		Spec: apiv1.VolumeSnapshotSpec{
			VolumeSnapshotClassName: &snapshotclassname,
			Source: apiv1.VolumeSnapshotSource{
				PersistentVolumeClaimName: &pvcname,
			},
		},
	}

	snapshotClient,err:=v1.NewForConfig(config)
	if err != nil {
	    log.Fatalln("failed to create snapshot client")
	}
	_,err=snapshotClient.VolumeSnapshots("default").Create(context.TODO(),volumeSnapshot,metav1.CreateOptions{})
	if err != nil {
	    log.Fatalln(err.Error())
	}

	
	// fileName := fmt.Sprintf("%s-%s-volume-snapshot.yaml", pvcname, namespace)
    // filePath := filepath.Join("backups", fileName)


    // err = os.WriteFile(filePath, depYAML, 0644)
    // if err != nil {
    //     panic(err.Error())
    // }
	
	// backup,_:=objects.CreateBackup(,namespace,"colume-snapshot","")
    // objects.AddBackup(*backup)
    // fmt.Println("deployment saved")

	
}

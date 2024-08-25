package internal

import (
    "fmt"
	"os"
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
    "gopkg.in/yaml.v3"
    "path/filepath"
    "K8Backup/objects"
	// clientcmd "k8s.io/client-go/tools/clientcmd"
)

func SavePod(clientset *kubernetes.Clientset,podName string,namespace string){

    podsClient := clientset.CoreV1().Pods(namespace)
    originalPod, Err := podsClient.Get(context.TODO(), podName, metav1.GetOptions{})

    if Err != nil {
        panic(Err.Error())
    }
   ///////////marshal to yaml

    podYAML, err := yaml.Marshal(originalPod)
    if err != nil {
        panic(err.Error())
    }


    // // // Write the YAML to a file

    fileName := fmt.Sprintf("%s-%s-pod.yaml", podName, namespace)
    filePath := filepath.Join("backups", fileName)

    // fileName := fmt.Sprintf("%s.yaml", podName)
    err = os.WriteFile(filePath, podYAML, 0644)
    if err != nil {
        panic(err.Error())
    }

    ////create backup obect

    backup,_:=objects.CreateBackup(podName,namespace,"pod",filePath)
    if(backup!=nil){
        fmt.Println("pod saved")
    }
    // objects.AddBackup(*backup)
    

    

    
    
	
}
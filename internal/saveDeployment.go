package internal

import (
    "fmt"
	"os"
	"context"
	// v1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
    "gopkg.in/yaml.v3"
    "path/filepath"
    "K8Backup/objects"
	// clientcmd "k8s.io/client-go/tools/clientcmd"
)

func SaveDeployment(clientset *kubernetes.Clientset,depName string,namespace string){

	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	dep,err:=deploymentsClient.Get(context.TODO(),depName,metav1.GetOptions{})
    if err!=nil{
        panic(err.Error())
    }

	
	depYAML, err := yaml.Marshal(dep)
    if err != nil {
        panic(err.Error())
    }

	fileName := fmt.Sprintf("%s-%s-dep.yaml", depName, namespace)
    filePath := filepath.Join("backups", fileName)


    err = os.WriteFile(filePath, depYAML, 0644)
    if err != nil {
        panic(err.Error())
    }
	
	backup,_:=objects.CreateBackup(depName,namespace,"deployment",filePath)
    if(backup!=nil){
        fmt.Println("deployment saved")
    }
    // objects.AddBackup(*backup)
   



    
    
	
}
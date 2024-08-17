package internal

import (
    "log"
    "fmt"
	"os"
	"context"
    "path/filepath"
    appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
    "gopkg.in/yaml.v3"
	// clientcmd "k8s.io/client-go/tools/clientcmd"
)

func RestoreDeployment(clientset *kubernetes.Clientset,fileName string) {
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	filePath := filepath.Join("backups", fileName)
    yamlData, err := os.ReadFile(filePath)
		if err != nil {
			panic(err.Error())
		}
    

    //////define new pod
    var newDep appsv1.Deployment
    err=yaml.Unmarshal(yamlData,&newDep)

    if err != nil {
        panic(err.Error())
    }
    log.Println("Done")
    

        deployment := &appsv1.Deployment{
            ObjectMeta: metav1.ObjectMeta{
                Name: "backedup-deployment-2",
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

        fmt.Println("Creating deployment...")
        _, err = deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
        if err != nil {
            panic(err)
        }
        
    
    
	
}
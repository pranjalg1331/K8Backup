package internal

import (
    "log"
    // "fmt"
	"os"
	"context"
    "path/filepath"
    v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
    "gopkg.in/yaml.v3"
	// clientcmd "k8s.io/client-go/tools/clientcmd"
)

func RestorePod(clientset *kubernetes.Clientset,fileName,restorename string) {
	filePath := filepath.Join(fileName)
    yamlData, err := os.ReadFile(filePath)
		if err != nil {
			panic(err.Error())
		}
    

    //////define new pod
    var newpod v1.Pod
    err=yaml.Unmarshal(yamlData,&newpod)

    if err != nil {
        panic(err.Error())
    }

    
    ///////config new pod
          new := &v1.Pod{
            ObjectMeta: metav1.ObjectMeta{
                Name:      restorename, // New name for the duplicate pod
                Namespace: newpod.Namespace,
                Labels:    newpod.Labels,
            },
            Spec: newpod.Spec,
        }

        
    /////run new pod
        _, err = clientset.CoreV1().Pods(newpod.Namespace).Create(context.TODO(), new, metav1.CreateOptions{})
    if err != nil {
        panic(err.Error())
    }
    log.Println("Creating Pod...")

    
    
	
}
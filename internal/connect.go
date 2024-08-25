package internal

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	kubernetes "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	clientcmd "k8s.io/client-go/tools/clientcmd"
)

func Connect(path string) *kubernetes.Clientset {
    
    config:=GetConfig(path)

    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        log.Panic("Failed to create K8s clientset")
    }
	
    return clientset
}

func GetConfig(path string) *rest.Config{
    fmt.Println(path)

    home,exists := os.LookupEnv("$HOME")
    if !exists {
        home = "C:\\Users\\Asush"
        // panic("dir not found")
    }
	log.Println(home)
    configPath := filepath.Join(home, ".kube", "config")
	log.Print(configPath)
    config, err := clientcmd.BuildConfigFromFlags("", path)
    if err != nil {
        log.Panic("failed to create K8s config")
    }

    return config
    
}
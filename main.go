package main

import (
	"fmt"
	"path/filepath"

	mygroup "github.com/Chandan-DK/kubernetes-custom-controller/pkg/generated/client/clientset/versioned"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type clients struct {
	kubernetesClient *kubernetes.Clientset
	mygroupClient    *mygroup.Clientset
}

func initClients() clients {
	// InitializeClients
	var c clients
	kubeconfigPath := filepath.Join(homedir.HomeDir(), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		fmt.Println(err)
	}
	c.kubernetesClient, err = kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println(err)
	}
	c.mygroupClient, err = mygroup.NewForConfig(config)
	if err != nil {
		fmt.Println(err)
	}
	return c
}
func main() {
	c := initClients()
	fmt.Println(c)
}

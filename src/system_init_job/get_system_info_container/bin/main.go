package main

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/Dat-Boi-Arjun/SEIFER/system_init_job/get_system_info_container/pkg"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func handle(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Running get_system_info main")
	var initNode = os.Getenv("INIT_NODE")
	ctx, cancel := context.WithCancel(context.Background())
	config, err := rest.InClusterConfig()
	handle(err)
	fmt.Println("Main authenticated in-cluster")

	clientset, err := kubernetes.NewForConfig(config)
	handle(err)

	list, err := clientset.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	handle(err)

	// Number of compute nodes
	// Before, would exclude the system init job node from the compute node list
	NumNodes := len(list.Items)
	fmt.Printf("Num compute nodes: %d\n", NumNodes)

	nodes := make([]string, 0, NumNodes)
	for _, item := range list.Items {
		nodes = append(nodes, item.Name)
	}

	connectionsToNode := system_info.GetConnectionsToNode(nodes)

	system_info.LaunchJobs(ctx, clientset, nodes, connectionsToNode)
	var wg sync.WaitGroup
	wg.Add(2)
	go system_info.Run(&wg, nodes, connectionsToNode)
	go system_info.ReceiveData(&wg, NumNodes)
	wg.Wait()

	system_info.DispatcherInitJob(ctx, clientset, initNode)

	// Stop any residual processes
	cancel()
}

package clientset

import (
	"context"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func Execute() {
	test()
}

func test() {
	// config
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}
	// client
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	// getdata
	pod, err := clientset.CoreV1().Pods("default").Get(context.TODO(), "test", v1.GetOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println(pod.Name)
}

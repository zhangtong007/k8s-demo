package dynamicclient

import (
	"context"
	"encoding/json"
	"os"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
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
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	// 数据处理
	// create resource
	resource := "pods"
	namespace := "default"
	pod := map[string]interface{}{
		"apiVersion": "v1",
		"kind":       "Pod",
		"metadata": map[string]interface{}{
			"name":      "mytest",
			"namespace": namespace,
		},
		"spec": map[string]interface{}{
			"containers": []map[string]interface{}{
				{
					"name":  "nginx",
					"image": "nginx:latest",
				},
			},
		},
	}
	r, err := dynamicClient.Resource(schema.GroupVersionResource{
		Group:    "",
		Version:  "v1",
		Resource: resource,
	}).Namespace(namespace).Create(context.TODO(), &unstructured.Unstructured{pod}, v1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	// query resource
	podName := r.GetName()
	getPod, err := dynamicClient.Resource(
		schema.GroupVersionResource{
			Group:    "",
			Version:  "v1",
			Resource: "pods",
		},
	).Namespace("default").Get(context.TODO(), podName, v1.GetOptions{})
	en := json.NewEncoder(os.Stdout)
	en.SetIndent("", "    ")
	en.Encode(getPod)
	// 删除资源
	err = dynamicClient.Resource(schema.GroupVersionResource{
		Group:    "",
		Version:  "v1",
		Resource: resource,
	}).Namespace(namespace).Delete(context.TODO(), podName, v1.DeleteOptions{})
	if err != nil {
		panic(err.Error())
	}
}

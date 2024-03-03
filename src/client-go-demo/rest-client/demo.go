package restclient

import (
	"context"
	"encoding/json"
	"os"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
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
	config.GroupVersion = &v1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs
	config.APIPath = "/api"
	// client
	rsClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err)
	}
	// get data
	pod := v1.Pod{}
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	err = rsClient.Get().Namespace("default").Resource("pods").Name("test").Do(ctx).Into(&pod)
	cancel()
	if err != nil {
		panic(err)
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "   ")
	enc.Encode(pod)
}

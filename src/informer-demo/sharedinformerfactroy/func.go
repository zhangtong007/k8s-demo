package sharedinformerfactroy

import (
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

func test() {
	// 1. 创建config
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}
	// 2. 创建client
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	// 3. informer factory
	//factory := informers.NewSharedInformerFactory(client, 0)
	// 可以通过withOptions 指定namespace订阅
	factory := informers.NewSharedInformerFactoryWithOptions(client, 0, informers.WithNamespace("default"))
	// 4. 绑定资源informer
	podInformer := factory.Core().V1().Pods().Informer()
	// 5. 绑定事件函数
	podInformer.AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc:    func(obj interface{}) {},
			UpdateFunc: func(oldObj, newObj interface{}) {},
			DeleteFunc: func(obj interface{}) {},
		},
	)
	// 6. 启动
	stopCh := make(chan struct{})
	factory.Start(stopCh)
	// 等待绑定的informers事件全部同步完成
	factory.WaitForCacheSync(stopCh)
	<-stopCh
}

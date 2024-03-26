package main

import (
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/workqueue"
	"time"
)

func main() {

	// config
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)

	if err != nil {
		panic(err.Error())
	}

	// client-set
	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		panic(err.Error())
	}
	informerFactory := informers.NewSharedInformerFactory(clientset, time.Minute)

	serviceInformer := informerFactory.Core().V1().Services()
	ingressInformer := informerFactory.Networking().V1().Ingresses()

	// 注册事件
	serviceInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				// AddFunc implementation
			},
			UpdateFunc: func(oldObj, newObj interface{}) {
				// UpdateFunc implementation
			},
			DeleteFunc: func(obj interface{}) {
				// DeleteFunc implementation
			},
		},
	)
	ingressInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				// AddFunc implementation
			},
			UpdateFunc: func(oldObj, newObj interface{}) {
				// UpdateFunc implementation
			},
			DeleteFunc: func(obj interface{}) {
				// DeleteFunc implementation
			},
		},
	)

	// 构造workerQueue，将事件推送到队列中
	workerQueue := workqueue.New()
}

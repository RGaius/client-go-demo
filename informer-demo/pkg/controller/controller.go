package controller

import (
	core "k8s.io/client-go/informers/core/v1"
	netInformer "k8s.io/client-go/informers/networking/v1"
	"k8s.io/client-go/kubernetes"
	netLister "k8s.io/client-go/listers/core/v1"
	coreLister "k8s.io/client-go/listers/networking/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

// Controller 构建结构体Controller
// client
// serviceLister
// ingressLister
// workerQueue
type Controller struct {
	client        kubernetes.Clientset
	serviceLister netLister.ServiceLister
	ingressLister coreLister.IngressLister
	workerQueue   workqueue.RateLimitingInterface
}

func (c *Controller) addService(obj interface{}) {

}

func (c *Controller) updateService(oldObject interface{}, newObj interface{}) {

}

func (c *Controller) deleteService(obj interface{}) {

}

// NewController 生成Controller构造函数
func NewController(client kubernetes.Clientset, serviceInformer core.ServiceInformer, ingressInformer netInformer.IngressInformer) *Controller {
	c := &Controller{
		client:        client,
		serviceLister: serviceInformer.Lister(),
		ingressLister: ingressInformer.Lister(),
		workerQueue:   workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter()),
	}

	serviceInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: c.addService, UpdateFunc: c.updateService})

	ingressInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{DeleteFunc: c.deleteService})
	return c
}

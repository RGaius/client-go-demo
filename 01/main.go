package main

import (
	"context"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// 读取配置文件
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err.Error())
	}
	// 构建client
	config.GroupVersion = &v1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs
	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err.Error())
	}
	pod := v1.Pod{}
	err = restClient.Get().Namespace("doau").Resource("pods").Name("deploy-xxljob-779d7579f5-rthsl").Do(context.TODO()).Into(&pod)
	if err != nil {
		println(err.Error())
	} else {
		println(pod.Name)
	}
}

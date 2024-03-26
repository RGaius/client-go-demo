package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// 家在配置文件
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)

	if err != nil {
		panic(err.Error())
	}
	// 构建client对象

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// 基于client-set获取pod对象
	_, err = clientset.CoreV1().Pods("doau").Get(context.TODO(), "deploy-xxljob-779d7579f5-rthsl", metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
		return
	}
	fmt.Printf("successful get pod %s in ns %s", "deploy-xxljob-779d7579f5-rthsl", "doau")
}

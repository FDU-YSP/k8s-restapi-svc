package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/golang/glog"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func main()  {
	flag.Set("alsologtostderr", "true")
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config-koor"), "(optional) ansolute path  to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

	glog.Info("list deploy")
	deployList, err := deploymentsClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		glog.Error("retrieve deploy list error")
	}
	for _, item := range deployList.Items {
		fmt.Printf(" **** %s (%d replicas)\n", item.Name, item.Spec.Replicas)
	}
}
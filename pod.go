package gcp

import (
	"context"
	"fmt"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func (g *Gcp) GetRunningPodByKeyword(namespace, keyword string) (string, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return "", fmt.Errorf("Can't load InCluster settings: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return "", fmt.Errorf("re-build clientset failure: %v", err)
	}

	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return "", fmt.Errorf("list Pod failure: %v", err)
	}

	for _, pod := range pods.Items {
		if pod.Status.Phase == "Running" && strings.Contains(pod.Name, keyword) {
			return pod.Name, nil
		}
	}

	return "", fmt.Errorf("can't find keyword [%s] and status is Running Pod", keyword)
}

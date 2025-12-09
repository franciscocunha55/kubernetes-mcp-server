package k8s


import (
	"context"
	"fmt"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ListNamespaces(clientSet *kubernetes.Clientset) *apiv1.NamespaceList{
	namespacesList, err := clientSet.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Found %d namespaces:\n", len(namespacesList.Items))
	for _, ns := range namespacesList.Items {
		fmt.Printf("- %s\n", ns.Name)
	}
	return namespacesList
}
package namespace

import (
	"context"
	"fmt"
	"strings"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// createNamespace creates a namespace with the given name and adds an annotation to it to enable Linkerd injection.
func CreateNamespace(clientset *kubernetes.Clientset, namespace string) error {
	fmt.Println("Creating namespace...")

	ns := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespace,
			Annotations: map[string]string{
				"linkerd.io/inject": "enabled",
			},
		},
	}

	_, err := clientset.CoreV1().Namespaces().Create(context.Background(), ns, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	fmt.Println("Namespace created.")
	return nil
}

func CheckIfNamespaceExists(clientset *kubernetes.Clientset, namespace string) (bool, error) {
	fmt.Println("Checking if namespace exists...")

	_, err := clientset.CoreV1().Namespaces().Get(context.Background(), namespace, metav1.GetOptions{})
	if err != nil {
		return false, err
	}

	fmt.Println("Namespace exists.")
	return true, nil
}

func DeleteNamespace(clientset *kubernetes.Clientset, namespace string) error {
	fmt.Println("Deleting namespace...")

	err := clientset.CoreV1().Namespaces().Delete(context.Background(), namespace, metav1.DeleteOptions{})
	if err != nil {
		return err
	}

	fmt.Println("Namespace deleted.")
	return nil
}

func GetListOfNamespacesWithPrefix(clientset *kubernetes.Clientset, prefix string) ([]string, error) {
	fmt.Println("Getting list of namespaces...")

	namespaces, err := clientset.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var listOfNamespaces []string
	for _, namespace := range namespaces.Items {
		if strings.HasPrefix(namespace.Name, prefix) {
			listOfNamespaces = append(listOfNamespaces, namespace.Name)
		}
	}

	fmt.Println("List of namespaces:", listOfNamespaces)
	return listOfNamespaces, nil
}

func GetListOfNamespacesWithSuffix(clientset *kubernetes.Clientset, suffix string) ([]string, error) {
	fmt.Println("Getting list of namespaces...")

	namespaces, err := clientset.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var listOfNamespaces []string
	for _, namespace := range namespaces.Items {
		if strings.HasSuffix(namespace.Name, suffix) {
			listOfNamespaces = append(listOfNamespaces, namespace.Name)
		}
	}

	fmt.Println("List of namespaces:", listOfNamespaces)
	return listOfNamespaces, nil
}

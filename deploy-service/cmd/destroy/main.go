package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/RobDoan/deploy-service/pkg/namespace"
	"github.com/RobDoan/deploy-service/pkg/utils"
)

type Options struct {
	JiraId string
}

func main() {

	var options Options

	flag.StringVar(&options.JiraId, "jira", "", "Jira ticket id")

	flag.Parse()

	k8sClient, err := utils.GetKubeClient()
	if err != nil {
		log.Fatalf("Failed to get client: %s", err)
	}

	namespaces, err := namespace.GetListOfNamespacesWithSuffix(k8sClient, fmt.Sprintf("-%s", options.JiraId))
	if err != nil {
		log.Fatalf("Failed to get list of namespaces: %s", err)
	}

	for _, ns := range namespaces {
		fmt.Printf("Deleting namespace %s\n", ns)
		err := namespace.DeleteNamespace(k8sClient, ns)
		if err != nil {
			log.Fatalf("Failed to delete namespace %s: %s", ns, err)
		}
	}
}

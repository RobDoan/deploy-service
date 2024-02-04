package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func GetServiceNameFromChart(chart string) string {
	split := strings.Split(chart, "/")
	var result string
	if len(split) > 1 && split[1] != "" {
		result = split[1]
	} else {
		result = split[0]
	}
	return result
}

func ExecuteCommand(command string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(output))
		panic(err)
	}
	fmt.Println(string(output))
	return string(output), nil
}

func GetKubeClient() (*kubernetes.Clientset, error) {
	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	k8sConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("Failed to build config: %s", err)
	}

	k8sClient, err := kubernetes.NewForConfig(k8sConfig)
	if err != nil {
		log.Fatalf("Failed to create client: %s", err)
	}
	return k8sClient, err
}

func GetJiraNumberFromNamespace(namespace string) string {
	re := regexp.MustCompile(`-jira-(\d+)$`)
	match := re.FindStringSubmatch(namespace)
	if len(match) < 2 {
		fmt.Println("Invalid namespace format")
		return ""
	}
	return "jira-" + match[1]
}

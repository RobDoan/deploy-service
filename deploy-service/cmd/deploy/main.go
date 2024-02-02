package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {

	var deployEngine = createDeployEngineFromCommandOptions()

	deployEngine.createNameSpace()

	deployEngine.createService()

	k8sClient, err := getKubeClient()
	if err != nil {
		log.Fatalf("Failed to get client: %s", err)
	}

	namespaces, err := GetListOfNamespacesWithPrefix(k8sClient, fmt.Sprintf("%s-", deployEngine.ServiceName))

	if err != nil {
		log.Fatalf("Failed to get list of namespaces: %s", err)
	}
	var rules = deployEngine.buildRouteRules(namespaces)

	routerBuilder := NewRouterBuilder(deployEngine.TemplatePath, deployEngine.ServiceName, deployEngine.ReleaseName, deployEngine.Port)

	httpRouter, err := routerBuilder.createHttpRouter(rules, deployEngine.Port)

	if err != nil {
		log.Fatalf("Failed to create http router: %s", err)
	}

	cmd := exec.Command("kubectl", "apply", "-f", "-")
	cmd.Stdin = bytes.NewBufferString(httpRouter)
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Failed to run kubectl: %v", err)
	}
	// Print the output of "kubectl apply"
	fmt.Println(out.String())

}

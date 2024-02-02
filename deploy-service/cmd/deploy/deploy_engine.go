package main

import (
	"bytes"
	"flag"
	"fmt"
	"regexp"
	"text/template"
)

// HelmCommandTpl represents helm command template
const HelmCommandTpl = "helm upgrade --install {{.ReleaseName}} -n {{.Namespace}} --set ui.message='{{.Namespace}} backend' {{.Chart}}"

// CommandOptions represents command line options

type DeployEngine struct {
	JiraId       string
	Chart        string
	ReleaseName  string
	ServiceName  string
	Namespace    string
	Port         int
	TemplatePath string
	Uat          bool
}

func createDeployEngineFromCommandOptions() *DeployEngine {
	var config DeployEngine
	flag.StringVar(&config.JiraId, "jira", "", "Jira ticket id")
	flag.StringVar(&config.Chart, "chart", "", "chart")
	flag.StringVar(&config.ReleaseName, "name", "", "release name")
	flag.StringVar(&config.ServiceName, "service", "", "Name of Service")
	flag.IntVar(&config.Port, "port", 9898, "Service Port")
	flag.BoolVar(&config.Uat, "uat", false, "UAT")

	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		config.TemplatePath = args[0]
	}

	if config.ServiceName == "" {
		config.ServiceName = getServiceNameFromChart(config.Chart)
	}

	if config.Uat {
		fmt.Println("Deploying UAT - JiraId will be ignored")
		config.JiraId = "uat"
	}

	if config.JiraId == "" {
		fmt.Println("JiraId is required")
		return nil
	}

	config.Namespace = fmt.Sprintf("%s-%s", config.ServiceName, config.JiraId)
	println(fmt.Sprintf("%s-%s", config.ServiceName, config.JiraId))
	return &config
}

func (de *DeployEngine) createNameSpace() {
	var createNamespaceCmd = fmt.Sprintf("kubectl create ns %s --dry-run=client -o yaml | linkerd inject -  | kubectl apply -f - \n", de.Namespace)
	fmt.Println(createNamespaceCmd)

	fmt.Println("Creating namespace...")
	fmt.Println(createNamespaceCmd)

	executeCommand(createNamespaceCmd)
}

func (de *DeployEngine) createService() {
	var buf bytes.Buffer

	t := template.Must(template.New("helmCommand").Parse("helm upgrade --install {{.ReleaseName}} -n {{.Namespace}} --set ui.message='{{.JiraId}} backend' {{.Chart}}\n"))
	t.Execute(&buf, de)

	fmt.Println(fmt.Sprintf("Creating helm chart: %s ...", de.Chart))
	hemlChartCmd := buf.String()
	executeCommand(hemlChartCmd)
}

func getJiraNumberFromNamespace(namespace string) string {
	re := regexp.MustCompile(`-jira-(\d+)$`)
	match := re.FindStringSubmatch(namespace)
	if len(match) < 2 {
		fmt.Println("Invalid namespace format")
		return ""
	}
	return "jira-" + match[1]
}

func isUATNamespace(namespace string) bool {
	re := regexp.MustCompile(`-uat$`)
	return re.MatchString(namespace)
}

func (de *DeployEngine) buildRouteRules(namespaces []string) []Rule {
	var rules []Rule
	for _, namespace := range namespaces {
		if isUATNamespace(namespace) {
			continue
		}
		rules = append(rules, createRule(namespace))
	}
	return rules
}

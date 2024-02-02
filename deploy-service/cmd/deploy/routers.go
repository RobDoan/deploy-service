package main

import (
	"bytes"
	"text/template"
)

type BackendRef struct {
	Name      string
	Namespace string
	Port      int
}

type Rule struct {
	RequestId string
	Namespace string
}

func createRule(namespace string) Rule {
	requestId := getJiraNumberFromNamespace(namespace)
	return Rule{
		RequestId: requestId,
		Namespace: namespace,
	}
}

type RouterBuilder struct {
	TemplatePath string
	ServiceName  string
	ReleaseName  string
	ServicePort  int
}

func NewRouterBuilder(path string, serviceName string, releaseName string, servicePort int) *RouterBuilder {
	return &RouterBuilder{
		TemplatePath: path,
		ServiceName:  serviceName,
		ReleaseName:  releaseName,
		ServicePort:  servicePort,
	}
}

// createHttpRouter creates a router)
func (b *RouterBuilder) createHttpRouter(rules []Rule, servicePort int) (string, error) {
	tmpl, err := template.ParseFiles(b.TemplatePath)

	if err != nil {
		return "", err
	}

	data := struct {
		ServiceName string
		ReleaseName string
		Rules       []Rule
		Port        int
	}{
		ServiceName: b.ServiceName,
		ReleaseName: b.ReleaseName,
		Rules:       rules,
		Port:        b.ServicePort,
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

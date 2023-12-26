package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	rust := sage{
		Name:  "Rust",
		Motto: "A language empowering everyone to build reliable and efficient software",
	}

	golang := sage{
		Name:  "Go",
		Motto: "Build simple, secure, scalable systems with Go",
	}

	docker := sage{
		Name:  "Docker",
		Motto: "Accelerate how you build, share, and run applications",
	}

	kubernetes := sage{
		Name:  "Kubernetes",
		Motto: "Kubernetes, also known as K8s, is an open-source system for automating deployment, scaling, and management of containerized applications",
	}

	aws := sage{
		Name:  "AWS",
		Motto: "AWS Cloud",
	}

	sages := []sage{rust, golang, docker, kubernetes, aws}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}

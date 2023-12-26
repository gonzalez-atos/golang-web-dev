package main

import (
	"fmt"
)

func main() {

	sages := map[string]string{
		"India":  "Docker",
		"Google": "Go",
		"GitHub": "Git",
		"K8s":    "Kubernetes",
		"Cloud":  "AWS",
	}

	for k, v := range sages {
		fmt.Println(k, " - ", v)
	}

}

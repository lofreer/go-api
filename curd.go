package main

import (
	"fmt"
	"os"

	"go-api/api/bootstrap"
	"go-api/api/pkg/generate"
)

func main() {
	model := os.Args[1]
	if model == "" {
		fmt.Println("参数错误")
	}

	caseVal := ""
	if len(os.Args) > 2 {
		caseVal = os.Args[2]
	}

	switch caseVal {
	case "route":
		generate.SetRoute(bootstrap.MigrateStruct[model])
	default:
		generate.SetCurd(bootstrap.MigrateStruct[model])
	}
}

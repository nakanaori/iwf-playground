package create_order

import (
	"fmt"
	"iwf-playground/repository"

	"github.com/indeedeng/iwf-golang-sdk/iwf"
)

func GetRegistry() (iwf.Registry, error) {
	registry := iwf.NewRegistry()

	err := registry.AddWorkflows(
		NewCreateOrderWorkflow(repository.NewOrderRepository()),
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return registry, nil
}

package models

import (
	"strings"

	"code.cloudfoundry.org/cli/plugin"
)

type ServiceInstance struct {
	GUID    string
	Name    string
	Service string
	Plan    string
}

func (si *ServiceInstance) ContainsTerms(items ...string) bool {
	for _, item := range items {
		item = strings.ToLower(item)
		service := strings.ToLower(si.Service)
		plan := strings.ToLower(si.Plan)
		if strings.Contains(service, item) || strings.Contains(plan, item) {
			return true
		}
	}
	return false
}

func FetchServiceInstance(cliConnection plugin.CliConnection, name string) (si ServiceInstance, err error) {
	srv, err := cliConnection.GetService(name)
	if err != nil {
		return
	}

	si = ServiceInstance{
		GUID:    srv.Guid,
		Service: srv.ServiceOffering.Name,
		Plan:    srv.ServicePlan.Name,
		Name:    name,
	}
	return
}

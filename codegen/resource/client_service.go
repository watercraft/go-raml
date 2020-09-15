package resource

import (
	"strings"

	"github.com/watercraft/go-raml/codegen/commons"
)

// ClientService defines a client service.
// client service is a module that handle
// a single root endpoint
type ClientService struct {
	Name          string
	EndpointName  string
	FilenameNoExt string
}

func NewClientService(rootEndpoint, displayName string) ClientService {
	var normalized string
	if displayName == "" {
		normalized = commons.NormalizeIdentifier(commons.NormalizeURI(rootEndpoint))
	} else {
		normalized = commons.DisplayNameToFuncName(displayName)
	}
	return ClientService{
		Name:          strings.Title(normalized) + "Service",
		EndpointName:  normalized,
		FilenameNoExt: normalized + "_service",
	}
}

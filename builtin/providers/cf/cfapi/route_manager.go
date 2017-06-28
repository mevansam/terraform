package cfapi

import (
	"code.cloudfoundry.org/cli/cf/api"
	"code.cloudfoundry.org/cli/cf/configuration/coreconfig"
	"code.cloudfoundry.org/cli/cf/net"
)

// RouteManager -
type RouteManager struct {
	log *Logger

	config    coreconfig.Reader
	ccGateway net.Gateway

	apiEndpoint string

	repo api.RouteRepository
}

// newRouteManager -
func newRouteManager(config coreconfig.Reader, ccGateway net.Gateway, logger *Logger) (rm *RouteManager, err error) {

	rm = &RouteManager{
		log: logger,

		config:    config,
		ccGateway: ccGateway,

		apiEndpoint: config.APIEndpoint(),

		repo: api.NewCloudControllerRouteRepository(config, ccGateway),
	}

	return
}

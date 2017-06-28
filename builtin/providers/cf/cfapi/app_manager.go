package cfapi

import (
	"code.cloudfoundry.org/cli/cf/api/applicationbits"
	"code.cloudfoundry.org/cli/cf/api/applications"
	"code.cloudfoundry.org/cli/cf/configuration/coreconfig"
	"code.cloudfoundry.org/cli/cf/net"
)

// AppManager -
type AppManager struct {
	log *Logger

	config    coreconfig.Reader
	ccGateway net.Gateway

	apiEndpoint string

	appRepo     applications.Repository
	appBitsRepo applicationbits.Repository
}

// newAppManager -
func newAppManager(config coreconfig.Reader, ccGateway net.Gateway, logger *Logger) (am *AppManager, err error) {

	am = &AppManager{
		log: logger,

		config:    config,
		ccGateway: ccGateway,

		apiEndpoint: config.APIEndpoint(),

		appRepo:     applications.NewCloudControllerRepository(config, ccGateway),
		appBitsRepo: applicationbits.NewCloudControllerApplicationBitsRepository(config, ccGateway),
	}

	return
}

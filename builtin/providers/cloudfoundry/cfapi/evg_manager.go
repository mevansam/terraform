package cfapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"code.cloudfoundry.org/cli/cf/configuration/coreconfig"
	"code.cloudfoundry.org/cli/cf/net"
)

// EVGManager -
type EVGManager struct {
	config    coreconfig.Reader
	ccGateway net.Gateway

	apiEndpoint string
}

// NewEVGManager -
func NewEVGManager(config coreconfig.Reader, ccGateway net.Gateway) (dm *EVGManager, err error) {

	dm = &EVGManager{
		config:    config,
		ccGateway: ccGateway,

		apiEndpoint: config.APIEndpoint(),
	}

	if dm.apiEndpoint == "" {
		err = errors.New("API endpoint missing from config file")
		return
	}

	return
}

// GetEVG -
func (dm *EVGManager) GetEVG(name string) (variables map[string]interface{}, err error) {

	url := fmt.Sprintf("%s/v2/config/environment_variable_groups/%s", dm.apiEndpoint, name)
	variables = make(map[string]interface{})
	err = dm.ccGateway.GetResource(url, &variables)
	return
}

// SetEVG -
func (dm *EVGManager) SetEVG(name string, variables map[string]interface{}) (err error) {

	body, err := json.Marshal(variables)
	if err != nil {
		return
	}

	err = dm.ccGateway.UpdateResource(dm.apiEndpoint,
		fmt.Sprintf("/v2/config/environment_variable_groups/%s", name),
		bytes.NewReader(body))

	return
}

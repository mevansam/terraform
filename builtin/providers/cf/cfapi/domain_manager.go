package cfapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"code.cloudfoundry.org/cli/cf/api"
	"code.cloudfoundry.org/cli/cf/api/resources"
	"code.cloudfoundry.org/cli/cf/configuration/coreconfig"
	"code.cloudfoundry.org/cli/cf/models"
	"code.cloudfoundry.org/cli/cf/net"
)

// DomainManager -
type DomainManager struct {
	log *Logger

	config    coreconfig.Reader
	ccGateway net.Gateway

	apiEndpoint string

	repo api.DomainRepository

	routingAPIRepo api.RoutingAPIRepository
}

// CCDomain -
type CCDomain struct {
	ID   string
	Name string `json:"name"`

	// Shared domain fields
	RouterGroupGUID string `json:"router_group_guid,omitempty"`
	RouterTypeGUID  string `json:"router_group_type,omitempty"`

	// Private domain fields
	OwningOrganizationGUID string `json:"owning_organization_guid,omitempty"`
}

// CCDomainResource -
type CCDomainResource struct {
	Metadata resources.Metadata `json:"metadata"`
	Entity   CCDomain           `json:"entity"`
}

// CCDomainList -
type CCDomainList struct {
	Resources []CCDomainResource `json:"resources"`
}

// NewDomainManager -
func newDomainManager(config coreconfig.Reader, ccGateway net.Gateway, logger *Logger) (dm *DomainManager, err error) {

	dm = &DomainManager{
		log: logger,

		config:    config,
		ccGateway: ccGateway,

		apiEndpoint: config.APIEndpoint(),

		repo:           api.NewCloudControllerDomainRepository(config, ccGateway),
		routingAPIRepo: api.NewRoutingAPIRepository(config, ccGateway),
	}

	if len(dm.apiEndpoint) == 0 {
		err = errors.New("API endpoint missing from config file")
		return
	}

	return
}

// GetSharedDomains -
func (dm *DomainManager) GetSharedDomains() (domains []CCDomain, err error) {

	domainList := CCDomainList{}
	err = dm.ccGateway.GetResource(fmt.Sprintf("%s/v2/shared_domains", dm.apiEndpoint), &domainList)

	for _, r := range domainList.Resources {
		domain := r.Entity
		domain.ID = r.Metadata.GUID
		domains = append(domains, domain)
	}
	return
}

// CreateSharedDomain -
func (dm *DomainManager) CreateSharedDomain(name string, routeGroupGUID *string) (domain CCDomain, err error) {

	var body []byte

	if routeGroupGUID != nil {
		body, err = json.Marshal(map[string]string{
			"name":              name,
			"router_group_guid": *routeGroupGUID,
		})
	} else {
		body, err = json.Marshal(map[string]string{
			"name": name,
		})
	}
	if err != nil {
		return
	}

	resource := CCDomainResource{}
	err = dm.ccGateway.CreateResource(dm.apiEndpoint, "/v2/shared_domains", bytes.NewReader(body), &resource)
	domain = resource.Entity
	domain.ID = resource.Metadata.GUID

	return
}

// GetSharedDomain -
func (dm *DomainManager) GetSharedDomain(guid string) (domain CCDomain, err error) {

	resource := CCDomainResource{}
	err = dm.ccGateway.GetResource(fmt.Sprintf("%s/v2/shared_domains/%s", dm.apiEndpoint, guid), &resource)
	domain = resource.Entity
	domain.ID = resource.Metadata.GUID

	return
}

// DeleteSharedDomain -
func (dm *DomainManager) DeleteSharedDomain(guid string) (err error) {
	err = dm.ccGateway.DeleteResource(dm.apiEndpoint, fmt.Sprintf("/v2/shared_domains/%s", guid))
	return
}

// GetPrivateDomains -
func (dm *DomainManager) GetPrivateDomains() (domains []CCDomain, err error) {

	domainList := CCDomainList{}
	err = dm.ccGateway.GetResource(fmt.Sprintf("%s/v2/private_domains", dm.apiEndpoint), &domainList)

	for _, r := range domainList.Resources {
		domain := r.Entity
		domain.ID = r.Metadata.GUID
		domains = append(domains, domain)
	}
	return
}

// CreatePrivateDomain -
func (dm *DomainManager) CreatePrivateDomain(name string, orgGUID string) (domain CCDomain, err error) {

	body, err := json.Marshal(map[string]string{
		"name": name,
		"owning_organization_guid": orgGUID,
	})
	if err != nil {
		return
	}

	resource := CCDomainResource{}
	err = dm.ccGateway.CreateResource(dm.apiEndpoint, "/v2/private_domains", bytes.NewReader(body), &resource)
	domain = resource.Entity
	domain.ID = resource.Metadata.GUID

	return
}

// GetPrivateDomain -
func (dm *DomainManager) GetPrivateDomain(guid string) (domain CCDomain, err error) {

	resource := &CCDomainResource{}
	err = dm.ccGateway.GetResource(fmt.Sprintf("%s/v2/private_domains/%s", dm.apiEndpoint, guid), resource)
	domain = resource.Entity
	domain.ID = resource.Metadata.GUID

	return
}

// DeletePrivateDomain -
func (dm *DomainManager) DeletePrivateDomain(guid string) (err error) {
	err = dm.ccGateway.DeleteResource(dm.apiEndpoint, fmt.Sprintf("/v2/private_domains/%s", guid))
	return
}

// FindSharedByName -
func (dm *DomainManager) FindSharedByName(name string) (models.DomainFields, error) {
	return dm.repo.FindSharedByName(name)
}

// FindPrivateByName -
func (dm *DomainManager) FindPrivateByName(name string) (models.DomainFields, error) {
	return dm.repo.FindPrivateByName(name)
}

// FindRouterGroupByName -
func (dm *DomainManager) FindRouterGroupByName(name string) (routerGroup models.RouterGroup, err error) {

	err = dm.routingAPIRepo.ListRouterGroups(
		func(rg models.RouterGroup) bool {
			if rg.Name == name {
				routerGroup = rg
				return false
			}
			return true
		})

	if routerGroup.Name != name {
		err = fmt.Errorf("Router group with name '%s' was not found.", name)
	}
	return
}

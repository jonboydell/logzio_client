package drop_filters

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client"
	"net/http"
)

const deactivateDropFilterServiceUrl string = dropFiltersServiceEndpoint + "/%s/deactivate"
const deactivateDropFilterServiceMethod string = http.MethodPost
const deactivateDropFilterMethodSuccess int = http.StatusOK
const deactivateDropFilterMethodNotFound int = http.StatusNotFound

func (c *DropFiltersClient) buildDeactivateApiRequest(apiToken string, dropFilterId string) (*http.Request, error) {
	baseUrl := c.BaseUrl
	req, err := http.NewRequest(deactivateDropFilterServiceMethod, fmt.Sprintf(deactivateDropFilterServiceUrl, baseUrl, dropFilterId), nil)
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

// Deactivates drop filter by Id, returns the deactivated drop filter or error if occurred
func (c *DropFiltersClient) DeactivateDropFilter(dropFilterId string) (*DropFilter, error) {
	return c.ActivateOrDeactivateDropFilter(dropFilterId, false)
}

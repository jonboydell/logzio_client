package drop_filters

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client"
	"net/http"
)

const activateDropFilterServiceUrl string = dropFiltersServiceEndpoint + "/%s/activate"
const activateDropFilterServiceMethod string = http.MethodPost
const activateDropFilterMethodSuccess int = http.StatusOK
const activateDropFilterMethodNotFound int = http.StatusNotFound

func (c *DropFiltersClient) buildActivateApiRequest(apiToken string, dropFilterId string) (*http.Request, error) {
	baseUrl := c.BaseUrl
	req, err := http.NewRequest(activateDropFilterServiceMethod, fmt.Sprintf(activateDropFilterServiceUrl, baseUrl, dropFilterId), nil)
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

// Activates drop filter by Id, returns the activated drop filter or error if occurred
func (c *DropFiltersClient) ActivateDropFilter(dropFilterID string) (*DropFilter, error) {
	return c.ActivateOrDeactivateDropFilter(dropFilterID, true)
}

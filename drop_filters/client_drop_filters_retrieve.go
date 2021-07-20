package drop_filters

import (
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client"
	"github.com/logzio/logzio_terraform_client/client"
	"io/ioutil"
	"net/http"
)

const retrieveDropFilterServiceUrl string = dropFiltersServiceEndpoint + "/search"
const retrieveDropFilterServiceMethod string = http.MethodPost
const retrieveDropFilterMethodSuccess int = http.StatusOK

func (c *DropFiltersClient) buildRetrieveApiRequest(apiToken string) (*http.Request, error) {
	baseUrl := c.BaseUrl
	req, err := http.NewRequest(retrieveDropFilterServiceMethod, fmt.Sprintf(retrieveDropFilterServiceUrl, baseUrl), nil)
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

// Returns all the drop filters in an array associated with the account identified by the supplied API token, returns an error if
// any problem occurs during the API call
func (c *DropFiltersClient) RetrieveDropFilters() ([]DropFilter, error) {
	req, _ := c.buildRetrieveApiRequest(c.ApiToken)

	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	jsonBytes, _ := ioutil.ReadAll(resp.Body)

	if !logzio_client.CheckValidStatus(resp, []int{retrieveDropFilterMethodSuccess}) {
		return nil, fmt.Errorf("API call %s failed with status code %d, data: %s", "RetrieveDropFilters", resp.StatusCode, jsonBytes)
	}

	var dropFilters []DropFilter
	err = json.Unmarshal(jsonBytes, &dropFilters)

	if err != nil {
		return nil, err
	}

	return dropFilters, nil
}

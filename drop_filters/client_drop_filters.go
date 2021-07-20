package drop_filters

import (
	"encoding/json"
	"fmt"
	logzio_client "github.com/logzio/logzio_terraform_client"
	"github.com/logzio/logzio_terraform_client/client"
	"io/ioutil"
	"net/http"
)

const dropFiltersServiceEndpoint string = "%s/v1/drop-filters"

type DropFiltersClient struct {
	*client.Client
}

type CreateDropFilter struct {
	LogType         string                 `json:"logType,omitempty"`
	FieldConditions []FieldConditionObject `json:"fieldConditions"`
}

type FieldConditionObject struct {
	FieldName string      `json:"fieldName"`
	Value     interface{} `json:"value"`
}

type DropFilter struct {
	Id             string                 `json:"id"`
	Active         bool                   `json:"active"`
	LogType        string                 `json:"logType"`
	FieldCondition []FieldConditionObject `json:"fieldConditions"`
}

func New(apiToken, baseUrl string) (*DropFiltersClient, error) {
	if len(apiToken) == 0 {
		return nil, fmt.Errorf("API token not defined")
	}
	if len(baseUrl) == 0 {
		return nil, fmt.Errorf("Base URL not defined")
	}
	c := &DropFiltersClient{
		Client: client.New(apiToken, baseUrl),
	}

	return c, nil
}

func validateCreateDropFilterRequest(filter CreateDropFilter) error {
	if len(filter.FieldConditions) == 0 {
		return fmt.Errorf("FieldConditions must be set")
	}

	for _, condition := range filter.FieldConditions {
		if len(condition.FieldName) == 0 {
			return fmt.Errorf("FieldName must be set")
		}

		if condition.Value == nil {
			return fmt.Errorf("Value must be set")
		}
	}

	return nil
}

// Activates/deactivates a drop filter given it's unique identifier. Returns the drop filter, an error otherwise
func (c *DropFiltersClient) ActivateOrDeactivateDropFilter(dropFilterId string, active bool) (*DropFilter, error) {
	var req *http.Request
	var operationName string
	if active {
		req, _ = c.buildActivateApiRequest(c.ApiToken, dropFilterId)
		operationName = "ActivateDropFilter"
	} else {
		req, _ = c.buildDeactivateApiRequest(c.ApiToken, dropFilterId)
		operationName = "DeactivateDropFilter"
	}

	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	jsonBytes, _ := ioutil.ReadAll(resp.Body)
	var dropFilter DropFilter
	err = json.Unmarshal(jsonBytes, &dropFilter)
	if err != nil {
		return nil, err
	}

	if !logzio_client.CheckValidStatus(resp, []int{activateDropFilterMethodSuccess, deactivateDropFilterMethodSuccess}) {
		if resp.StatusCode == activateDropFilterMethodNotFound || resp.StatusCode == deactivateDropFilterMethodNotFound {
			return nil, fmt.Errorf("API call %s failed with missing drop filter %s, data: %s", operationName, dropFilterId, jsonBytes)
		}

		return nil, fmt.Errorf("API call %s failed with status code %d, data: %s", operationName, resp.StatusCode, jsonBytes)
	}

	return &dropFilter, nil
}

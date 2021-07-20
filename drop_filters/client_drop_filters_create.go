package drop_filters

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/logzio/logzio_terraform_client"
	"net/http"
)

const createDropFilterServiceUrl = dropFiltersServiceEndpoint
const createDropFilterServiceMethod string = http.MethodPost

type FieldError struct {
	Field   string
	Message string
}

func (e FieldError) Error() string {
	return fmt.Sprintf("%v: %v", e.Field, e.Message)
}

// Create a drop filter, return the created drop filter if successful, an error otherwise
func (c *DropFiltersClient) CreateDropFilter(createDropFilter CreateDropFilter) (*DropFilter, error) {
	err := validateCreateDropFilterRequest(createDropFilter)
	if err != nil {
		return nil, err
	}

	dropFilterJson, err := json.Marshal(createDropFilter)
	if err != nil {
		return nil, err
	}

	req, err := c.buildCreateApiRequest(c.ApiToken, dropFilterJson)
	if err != nil {
		return nil, err
	}

	jsonResponse, err := logzio_client.CreateHttpRequestBytesResponse(req)
	if err != nil {
		return nil, err
	}

	var retVal DropFilter
	err = json.Unmarshal(jsonResponse, &retVal)
	if err != nil {
		return nil, err
	}

	return &retVal, nil

}

func (c *DropFiltersClient) buildCreateApiRequest(apiToken string, jsonBytes []byte) (*http.Request, error) {
	baseUrl := c.BaseUrl
	req, err := http.NewRequest(createDropFilterServiceMethod, fmt.Sprintf(createDropFilterServiceUrl, baseUrl), bytes.NewBuffer(jsonBytes))
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

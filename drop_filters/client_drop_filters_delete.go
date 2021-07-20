package drop_filters

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client"
	"github.com/logzio/logzio_terraform_client/client"
	"io/ioutil"
	"net/http"
)

const deleteDropFilterServiceUrl string = dropFiltersServiceEndpoint + "/%s"
const deleteDropFilterServiceMethod string = http.MethodDelete
const deleteDropFilterMethodSuccess int = http.StatusOK
const deleteDropFilterMethodNotFound int = http.StatusNotFound

func (c *DropFiltersClient) buildDeleteApiRequest(apiToken string, dropFilterId string) (*http.Request, error) {
	baseUrl := c.BaseUrl
	req, err := http.NewRequest(deleteDropFilterServiceMethod, fmt.Sprintf(deleteDropFilterServiceUrl, baseUrl, dropFilterId), nil)
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

// Delete a drop filter, specified by it's unique id, returns an error if a problem is encountered
func (c *DropFiltersClient) DeleteDropFilter(dropFilterId string) error {
	req, _ := c.buildDeleteApiRequest(c.ApiToken, dropFilterId)
	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	jsonBytes, _ := ioutil.ReadAll(resp.Body)

	if !logzio_client.CheckValidStatus(resp, []int{deleteDropFilterMethodSuccess}) {
		if resp.StatusCode == deleteDropFilterMethodNotFound {
			return fmt.Errorf("API call %s failed with missing drop filter %s, data: %s", "DeleteDropFilter", dropFilterId, jsonBytes)
		}

		return fmt.Errorf("API call %s failed with status code %d, data: %s", "DeleteDropFilter", resp.StatusCode, jsonBytes)
	}

	return nil
}

package logshipping

import (
	"fmt"

	"github.com/logzio/logzio_terraform_client/client"
)

const (
	serviceEndpoint = "%s/v1/log-shipping"

	fldID       string = "id"
	fldLogsType string = "logsType"
)

// Client is the client used to talk to the logzio log shipping API
type Client struct {
	*client.Client
}

// New creates an entry point into the logshipping functions. It accepts the user's logz.io API token and account Id
func New(apiToken, baseURL string) (*Client, error) {
	if len(apiToken) == 0 {
		return nil, fmt.Errorf("API token not defined")
	}
	if len(baseUrl) == 0 {
		return nil, fmt.Errorf("Base URL not defined")
	}
	c := &LogShippingClient{
		Client: client.New(apiToken, baseUrl),
	}
	return c, nil
}

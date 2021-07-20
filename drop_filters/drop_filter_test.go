package drop_filters_test

import (
	"github.com/logzio/logzio_terraform_client/drop_filters"
	"github.com/logzio/logzio_terraform_client/test_utils"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
)

func fixture(path string) string {
	b, err := ioutil.ReadFile("testdata/fixtures/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func setupDropFiltersTest() (*drop_filters.DropFiltersClient, error, func()) {
	apiToken := "SOME_API_TOKEN"

	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	underTest, _ := drop_filters.New(apiToken, server.URL)

	return underTest, nil, func() {
		server.Close()
	}
}

func setupDropFiltersIntegrationTest() (*drop_filters.DropFiltersClient, error) {
	apiToken, err := test_utils.GetApiToken()
	if err != nil {
		return nil, err
	}

	underTest, err := drop_filters.New(apiToken, test_utils.GetLogzIoBaseUrl())
	return underTest, nil
}

func getCreateDropFilter() drop_filters.CreateDropFilter {
	fieldCondition := drop_filters.FieldConditionObject{
		FieldName: "some_field",
		Value:     "ok",
	}

	return drop_filters.CreateDropFilter{
		LogType:         "some_type",
		FieldConditions: []drop_filters.FieldConditionObject{fieldCondition},
	}
}
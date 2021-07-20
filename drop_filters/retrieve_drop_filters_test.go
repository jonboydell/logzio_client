package drop_filters_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestDropFilters_RetrieveDropFilter(t *testing.T) {
	underTest, err, teardown := setupDropFiltersTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/drop-filters/search", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("retrieve_drop_filter.json"))
		})

		filters, err := underTest.RetrieveDropFilters()
		assert.NoError(t, err)
		assert.NotNil(t, filters)
		assert.Equal(t, 2, len(filters))
	}
}

func TestDropFilters_RetrieveDropFilterAPIFail(t *testing.T) {
	underTest, err, teardown := setupDropFiltersTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/drop-filters/search", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, fixture("retrieve_drop_filter_failed.txt"))
		})

		filters, err := underTest.RetrieveDropFilters()
		assert.Error(t, err)
		assert.Nil(t, filters)
	}
}

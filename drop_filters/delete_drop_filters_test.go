package drop_filters_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestDropFilters_DeleteDropFilter(t *testing.T) {
	underTest, err, teardown := setupDropFiltersTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/drop-filters/", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodDelete, r.Method)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, fixture("delete_drop_filter.json"))
		})

		id := "some-drop-filter-id"
		err = underTest.DeleteDropFilter(id)
		assert.NoError(t, err)
	}
}

func TestDropFilters_DeleteDropFilterAPIFail(t *testing.T) {
	underTest, err, teardown := setupDropFiltersTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/drop-filters/", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodDelete, r.Method)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, fixture("delete_drop_filter_failed.txt"))
		})

		id := "some-drop-filter-id"
		err = underTest.DeleteDropFilter(id)
		assert.Error(t, err)
	}
}

func TestDropFilters_DeleteDropFilterNotFound(t *testing.T) {
	underTest, err, teardown := setupDropFiltersTest()
	defer teardown()

	if assert.NoError(t, err) {
		mux.HandleFunc("/v1/drop-filters/", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodDelete, r.Method)
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, fixture("delete_drop_filter_not_found.txt"))
		})

		id := "some-drop-filter-id-not-exist"
		err = underTest.DeleteDropFilter(id)
		assert.Error(t, err)
	}
}

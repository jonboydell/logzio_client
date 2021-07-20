package drop_filters_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationDropFilters_RetrieveDropFilter(t *testing.T) {
	underTest, err := setupDropFiltersIntegrationTest()

	if assert.NoError(t, err) {
		filters, err := underTest.RetrieveDropFilters()
		assert.NoError(t, err)
		assert.NotNil(t, filters)
	}
}

func TestIntegrationDropFilters_RetrieveDropFilterWithAtLeastOne(t *testing.T) {
	underTest, err := setupDropFiltersIntegrationTest()

	if assert.NoError(t, err) {
		createDropFilter := getCreateDropFilter()
		// Create a drop filter to make sure there's at least one
		dropFilter, err := underTest.CreateDropFilter(createDropFilter)

		time.Sleep(2 * time.Second)
		if assert.NoError(t, err) && assert.NotNil(t, dropFilter) {
			defer underTest.DeleteDropFilter(dropFilter.Id)
			filters, err := underTest.RetrieveDropFilters()
			assert.NoError(t, err)
			assert.NotNil(t, filters)
		}
	}
}

package drop_filters_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationDropFilters_DeleteDropFilter(t *testing.T) {
	underTest, err := setupDropFiltersIntegrationTest()

	if assert.NoError(t, err) {
		createDropFilter := getCreateDropFilter()
		dropFilter, err := underTest.CreateDropFilter(createDropFilter)

		if assert.NoError(t, err) && assert.NotNil(t, dropFilter) {
			time.Sleep(2 * time.Second)
			defer func() {
				err = underTest.DeleteDropFilter(dropFilter.Id)
				assert.NoError(t, err)
			}()
		}
	}
}

func TestIntegrationDropFilters_DeleteDropFilterNotFound(t *testing.T) {
	underTest, err := setupDropFiltersIntegrationTest()

	if assert.NoError(t, err) {
		err = underTest.DeleteDropFilter("id-not-exist")
		assert.Error(t, err)
	}
}

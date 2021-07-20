package drop_filters_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationDropFilters_ActivateDropFilter(t *testing.T) {
	underTest, err := setupDropFiltersIntegrationTest()

	if assert.NoError(t, err) {
		createDropFilter := getCreateDropFilter()
		dropFilter, err := underTest.CreateDropFilter(createDropFilter)

		if assert.NoError(t, err) && assert.NotNil(t, dropFilter) {
			time.Sleep(2 * time.Second)
			defer underTest.DeleteDropFilter(dropFilter.Id)
			_, err = underTest.ActivateOrDeactivateDropFilter(dropFilter.Id, false)
			if assert.NoError(t, err) {
				time.Sleep(2 * time.Second)
				activated, err := underTest.ActivateDropFilter(dropFilter.Id)
				assert.NoError(t, err)
				assert.NotNil(t, activated)
				assert.True(t, activated.Active)
			}
		}
	}
}

func TestIntegrationDropFilters_ActivateDropFilterNotFound(t *testing.T) {
	underTest, err := setupDropFiltersIntegrationTest()

	if assert.NoError(t, err) {
		id := "some-invalid-id"
		dropFilter, err := underTest.ActivateDropFilter(id)
		assert.Error(t, err)
		assert.Nil(t, dropFilter)
	}
}

func TestIntegrationDropFilters_ActivateDropFilterNoId(t *testing.T) {
	underTest, err := setupDropFiltersIntegrationTest()

	if assert.NoError(t, err) {
		dropFilter, err := underTest.ActivateDropFilter("")
		assert.Error(t, err)
		assert.Nil(t, dropFilter)
	}
}

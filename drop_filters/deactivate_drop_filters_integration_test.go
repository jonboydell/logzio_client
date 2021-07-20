package drop_filters_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationDropFilters_DeactivateDropFilter(t *testing.T) {
	underTest, err := setupDropFiltersIntegrationTest()

	if assert.NoError(t, err) {
		createDropFilter := getCreateDropFilter()
		dropFilter, err := underTest.CreateDropFilter(createDropFilter)

		if assert.NoError(t, err) && assert.NotNil(t, dropFilter) {
			time.Sleep(2 * time.Second)
			defer underTest.DeleteDropFilter(dropFilter.Id)
			deactivated, err := underTest.DeactivateDropFilter(dropFilter.Id)
			assert.NoError(t, err)
			assert.NotNil(t, deactivated)
			assert.False(t, deactivated.Active)

		}
	}
}

func TestIntegrationDropFilters_DeactivateDropFilterNotFound(t *testing.T) {
	underTest, err := setupDropFiltersIntegrationTest()

	if assert.NoError(t, err) {
		id := "some-invalid-id"
		dropFilter, err := underTest.DeactivateDropFilter(id)
		assert.Error(t, err)
		assert.Nil(t, dropFilter)
	}
}

func TestIntegrationDropFilters_DeactivateDropFilterNoId(t *testing.T) {
	underTest, err := setupDropFiltersIntegrationTest()

	if assert.NoError(t, err) {
		dropFilter, err := underTest.DeactivateDropFilter("")
		assert.Error(t, err)
		assert.Nil(t, dropFilter)
	}
}

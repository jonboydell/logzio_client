package drop_filters_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationDropFilters_CreateDropFilter(t *testing.T) {
	underTest, err := setupDropFiltersIntegrationTest()

	if assert.NoError(t, err) {
		createDropFilter := getCreateDropFilter()
		dropFilter, err := underTest.CreateDropFilter(createDropFilter)

		time.Sleep(2 * time.Second)
		if assert.NoError(t, err) && assert.NotNil(t, dropFilter) {
			defer underTest.DeleteDropFilter(dropFilter.Id)
		}
	}
}

func TestIntegrationDropFilters_CreateDropFilterNoLogType(t *testing.T) {
	underTest, err := setupDropFiltersIntegrationTest()

	if assert.NoError(t, err) {
		createDropFilter := getCreateDropFilter()
		createDropFilter.LogType = ""
		dropFilter, err := underTest.CreateDropFilter(createDropFilter)

		time.Sleep(2 * time.Second)
		if assert.NoError(t, err) && assert.NotNil(t, dropFilter) {
			defer underTest.DeleteDropFilter(dropFilter.Id)
		}
	}
}

func TestIntegrationDropFilters_CreateDropFilterNoConditions(t *testing.T) {
	underTest, err := setupDropFiltersIntegrationTest()

	if assert.NoError(t, err) {
		createDropFilter := getCreateDropFilter()
		createDropFilter.FieldConditions = nil
		dropFilter, err := underTest.CreateDropFilter(createDropFilter)
		assert.Error(t, err)
		assert.Nil(t, dropFilter)
	}
}

func TestIntegrationDropFilters_CreateDropFilterNoConditionFieldName(t *testing.T) {
	underTest, err := setupDropFiltersIntegrationTest()

	if assert.NoError(t, err) {
		createDropFilter := getCreateDropFilter()
		createDropFilter.FieldConditions[0].FieldName = ""
		dropFilter, err := underTest.CreateDropFilter(createDropFilter)
		assert.Error(t, err)
		assert.Nil(t, dropFilter)
	}
}

func TestIntegrationDropFilters_CreateDropFilterNoConditionValue(t *testing.T) {
	underTest, err := setupDropFiltersIntegrationTest()

	if assert.NoError(t, err) {
		createDropFilter := getCreateDropFilter()
		createDropFilter.FieldConditions[0].Value = nil
		dropFilter, err := underTest.CreateDropFilter(createDropFilter)
		assert.Error(t, err)
		assert.Nil(t, dropFilter)
	}
}

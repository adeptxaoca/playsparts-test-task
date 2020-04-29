package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	v := New()
	require.NotNil(t, v)
}

func TestValidator_Struct(t *testing.T) {
	type item struct {
		Name   string `validate:"name"`
		Code   string `validate:"vendor-code"`
		Number int32  `validate:"gt=100"`
	}

	v := New()
	err := v.Struct(&item{Name: "item_1", Code: "123456", Number: 50})
	if assert.Error(t, err) {
		assert.Equal(t, "item.Name,item.Number", err.Error())
	}
	err = v.Struct(&item{Name: "item_2", Code: "#123", Number: 10})
	if assert.Error(t, err) {
		assert.Equal(t, "item.Name,item.Code,item.Number", err.Error())
	}
	err = v.Struct(&item{Name: "item 2", Code: "ASD123", Number: 110})
	assert.NoError(t, err)
}

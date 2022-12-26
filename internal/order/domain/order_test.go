package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenEmpyId_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{}
	assert.Error(t, order.IsValid(), "invalid id")
}

func TestGivenEmpyPrice_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{ID: "123"}
	assert.Error(t, order.IsValid(), "invalid price")
}

func TestGivenEmpyTax_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{ID: "123", Price: 10.0}
	assert.Error(t, order.IsValid(), "invalid tax")
}

func TestGivenAValidParams_WhenICallNewOrder_ThenIShouldReceiveCreateOrderWithAllParams(t *testing.T) {
	order := Order{ID: "123", Price: 10.0, Tax: 1.0}
	assert.NoError(t, order.IsValid())
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 1.0, order.Tax)
	assert.Nil(t, order.IsValid())
}

func TestGivenAValidParams_WhenICallNewOrderFunc_ThenIShouldReceiveCreateOrderWithAllParams(t *testing.T) {
	order, err := NewOrder("123", 10.0, 1.0)
	assert.Nil(t, err)
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 1.0, order.Tax)
}
func TestGivenAPriceAndTax_WhenCallCalculatePrice_ThenIShoudSetFinalPrice(t *testing.T) {
	order, err := NewOrder("123", 10.0, 1.0)
	assert.Nil(t, err)
	assert.Nil(t, order.CalculatePrice())
	assert.Equal(t, 11.0, order.FinalPrice)
}

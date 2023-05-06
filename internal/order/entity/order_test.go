package entity_test

import (
	"github.com/joseMarciano/intensive-golang/internal/order/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGivenAnEmptyId_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := entity.Order{}
	assert.Error(t, order.IsValid(), "invalid id")
}

func TestGivenAnEmptyPrice_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := entity.Order{ID: "123"}
	assert.Error(t, order.IsValid(), "invalid price")
}

func TestGivenAnEmptyTax_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := entity.Order{ID: "123", Price: 10}
	assert.Error(t, order.IsValid(), "invalid tax")
}

func TestGivenAValidParams_WhenCallNewOrder_ThenShouldReceiveCreateOrderWithALlParams(t *testing.T) {
	order, err := entity.NewOrder("123", 10.5, 2)
	assert.NoError(t, err)
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.5, order.Price)
	assert.Equal(t, float64(2), order.Tax)
}

func TestGivenAValidParams_WhenCallCalculateFinalPrice_ThenShouldCalculateFinalPriceAndSetItOnFinalPriceProperty(t *testing.T) {
	order, err := entity.NewOrder("123", 10.5, 2)
	assert.NoError(t, err)
	err = order.CalculateFinalPrice()
	assert.NoError(t, err)
	assert.Equal(t, 12.5, order.FinalPrice)
}

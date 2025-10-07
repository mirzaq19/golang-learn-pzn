package service

import (
	"golang-unit-test/entity"
	"golang-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var itemRepository = &repository.ItemRepositoryMock{Mock: mock.Mock{}}
var itemService = ItemService{Repository: itemRepository}

func TestItemService_GetNotFound(t *testing.T) {
	itemRepository.Mock.On("FindById", "1").Return(nil)
	item, err := itemService.Get("1")
	itemRepository.Mock.AssertExpectations(t)
	assert.NotNil(t, err)
	assert.Equal(t, "item not found", err.Error(), "Result must be 'item not found'")
	assert.Nil(t, item)
}

func TestItemService_GetSuccess(t *testing.T) {
	item := entity.Item{
		Id:    "2",
		Name:  "HP Pavilion Gaming 15",
		Price: 14000000,
	}
	itemRepository.Mock.On("FindById", "2").Return(item)
	result, err := itemService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, item.Id, result.Id)
	assert.Equal(t, item.Name, result.Name)
	assert.Equal(t, item.Price, result.Price)
}

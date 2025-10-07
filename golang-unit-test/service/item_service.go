package service

import (
	"errors"
	"golang-unit-test/entity"
	"golang-unit-test/repository"
)

type ItemService struct {
	Repository repository.ItemRepository
}

func (service *ItemService) Get(id string) (*entity.Item, error) {
	item := service.Repository.FindById(id)
	if item == nil {
		return nil, errors.New("item not found")
	} else {
		return item, nil
	}
}

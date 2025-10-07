package repository

import (
	"golang-unit-test/entity"

	"github.com/stretchr/testify/mock"
)

type ItemRepositoryMock struct {
	Mock mock.Mock
}

func (repository *ItemRepositoryMock) FindById(id string) *entity.Item {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	}

	item := arguments.Get(0).(entity.Item)
	return &item
}

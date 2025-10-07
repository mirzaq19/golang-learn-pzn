package repository

import "golang-unit-test/entity"

type ItemRepository interface {
	FindById(id string) *entity.Item
}

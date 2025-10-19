package simple

import "errors"

type SimpleRepository struct {
	Error bool
}

func NewSimpleRepository(isError bool) *SimpleRepository {
	return &SimpleRepository{
		Error: isError,
	}
}

type SimpleService struct {
	*SimpleRepository
}

func NewSimpleService(r *SimpleRepository) (*SimpleService, error) {
	if r.Error {
		return nil, errors.New("failed to create service")
	} else {
		return &SimpleService{
			SimpleRepository: r,
		}, nil
	}
}

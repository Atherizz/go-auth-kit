package simple

import "errors"

type SampleRepository struct {
    Error bool
}

func NewSampleRepository(isError bool) *SampleRepository {
    return &SampleRepository{
        Error: isError,
    }
}

type SampleService struct {
    *SampleRepository
}

func NewSampleService(repository *SampleRepository) (*SampleService, error) {
    if repository.Error {
        return nil, errors.New("Failed create service")
    } else {
    return &SampleService{
        SampleRepository: repository,
    }, nil
    }
}

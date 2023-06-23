package impl

import (
	"Technopark_DB_Project/internal/models"
	"Technopark_DB_Project/internal/repositories"
	"Technopark_DB_Project/internal/usecases"
)

type ServiceUseCaseImpl struct {
	serviceRepository repositories.ServiceRepository
}

func CreateServiceUseCase(serviceRepository repositories.ServiceRepository) usecases.ServiceUseCase {
	return &ServiceUseCaseImpl{serviceRepository: serviceRepository}
}

func (serviceUseCase *ServiceUseCaseImpl) Clear() (err error) {
	return serviceUseCase.serviceRepository.Clear()
}

func (serviceUseCase *ServiceUseCaseImpl) GetStatus() (status *models.Status, err error) {
	return serviceUseCase.serviceRepository.GetStatus()
}

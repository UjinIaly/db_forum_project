package usecases

import "Technopark_DB_Project/internal/models"

type ServiceUseCase interface {
	Clear() (err error)
	GetStatus() (status *models.Status, err error)
}

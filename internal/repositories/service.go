package repositories

import "Technopark_DB_Project/internal/models"

type ServiceRepository interface {
	Clear() (err error)
	GetStatus() (status *models.Status, err error)
}

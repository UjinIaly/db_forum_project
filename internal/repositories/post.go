package repositories

import "Technopark_DB_Project/internal/models"

type PostRepository interface {
	GetByID(id int64) (post *models.Post, err error)
	Update(post *models.Post) (err error)
}

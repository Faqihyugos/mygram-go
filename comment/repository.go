package comment

import "gorm.io/gorm"

type Repository interface {
	Create(comment Comment) (Comment, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(comment Comment) (Comment, error) {
	err := r.db.Create(&comment).Error
	if err != nil {
		return comment, err
	}
	return comment, nil
}

package sosmed

import "gorm.io/gorm"

type Repository interface {
	Create(sosmed Sosmed) (Sosmed, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(sosmed Sosmed) (Sosmed, error) {
	err := r.db.Create(&sosmed).Error
	if err != nil {
		return sosmed, err
	}
	return sosmed, nil
}

package sosmed

import "gorm.io/gorm"

type Repository interface {
	Create(sosmed Sosmed) (Sosmed, error)
	FindAll() ([]Sosmed, error)
	Update(sosmed Sosmed) (Sosmed, error)
	FindByID(ID int) (Sosmed, error)
	Delete(sosmed Sosmed) (Sosmed, error)
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

func (r *repository) FindAll() ([]Sosmed, error) {
	var socialmedias []Sosmed
	err := r.db.Preload("User").Order("id asc").Find(&socialmedias).Error
	if err != nil {
		return socialmedias, err
	}
	return socialmedias, nil
}

func (r *repository) Update(sosmed Sosmed) (Sosmed, error) {
	err := r.db.Save(&sosmed).Error
	if err != nil {
		return sosmed, err
	}
	return sosmed, nil
}

func (r *repository) FindByID(ID int) (Sosmed, error) {
	var sosmed Sosmed
	err := r.db.Preload("User").Where("id = ?", ID).Find(&sosmed).Error
	if err != nil {
		return sosmed, err
	}
	return sosmed, nil
}

func (r *repository) Delete(sosmed Sosmed) (Sosmed, error) {
	err := r.db.Delete(&sosmed).Error
	if err != nil {
		return sosmed, err
	}
	return sosmed, nil
}

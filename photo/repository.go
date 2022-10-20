package photo

import "gorm.io/gorm"

type Repository interface {
	Save(photo Photo) (Photo, error)
	FindAll() ([]Photo, error)
	// FindByEmail(email string) (Photo, error)
	// FindByID(ID int) (Photo, error)
	// Update(photo Photo) (Photo, error)
	// Delete(photo Photo) (Photo, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(photo Photo) (Photo, error) {
	err := r.db.Create(&photo).Error
	if err != nil {
		return photo, err
	}
	return photo, nil
}

func (r *repository) FindAll() ([]Photo, error) {
	var photos []Photo
	err := r.db.Preload("User").Order("id asc").Find(&photos).Error
	if err != nil {
		return photos, err
	}
	return photos, nil
}

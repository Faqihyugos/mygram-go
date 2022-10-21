package comment

import "gorm.io/gorm"

type Repository interface {
	Create(comment Comment) (Comment, error)
	FindAll() ([]Comment, error)
	Update(comment Comment) (Comment, error)
	FindByID(ID int) (Comment, error)
	Delete(comment Comment) (Comment, error)
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

func (r *repository) FindAll() ([]Comment, error) {
	var comments []Comment
	err := r.db.Preload("User").Preload("Photo").Order("id asc").Find(&comments).Error
	if err != nil {
		return comments, err
	}
	return comments, nil
}

func (r *repository) Update(comment Comment) (Comment, error) {
	err := r.db.Save(&comment).Error
	if err != nil {
		return comment, err
	}
	return comment, nil
}

func (r *repository) FindByID(ID int) (Comment, error) {
	var comment Comment
	err := r.db.Preload("User").Preload("Photo").Where("id = ?", ID).Find(&comment).Error
	if err != nil {
		return comment, err
	}
	return comment, nil
}

func (r *repository) Delete(comment Comment) (Comment, error) {
	err := r.db.Delete(&comment).Error
	if err != nil {
		return comment, err
	}
	return comment, nil
}

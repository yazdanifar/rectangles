package repository

import (
	"gorm.io/gorm"
	"rectangles/model"
)

type RectangleRepository struct {
	db *gorm.DB
}

func NewRectangleRepository(db *gorm.DB) *RectangleRepository {
	return &RectangleRepository{db: db}
}

func (r *RectangleRepository) Create(rectangle *model.Rectangle) {
	r.db.Create(rectangle)
}

func (r *RectangleRepository) GetAllRectangles() ([]model.Rectangle, error) {
	var rectangles []model.Rectangle
	if err := r.db.Find(&rectangles).Error; err != nil {
		return nil, err
	}
	return rectangles, nil
}

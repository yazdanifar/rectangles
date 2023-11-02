package service

import (
	"rectangles/model"
	"rectangles/repository"
)

type RectangleService struct {
	repo *repository.RectangleRepository
}

func NewRectangleService(repo *repository.RectangleRepository) *RectangleService {
	return &RectangleService{repo: repo}
}

func (s *RectangleService) CreateIntersectingRectangles(main model.Rectangle, input []model.Rectangle) []model.Rectangle {

	intersectingRectangles := []model.Rectangle{}
	for _, rectangle := range input {
		if main.Intersects(rectangle) {
			intersectingRectangles = append(intersectingRectangles, rectangle)
			s.repo.Create(&rectangle)
		}
	}
	return intersectingRectangles
}

func (s *RectangleService) GetAllRectangles() ([]model.Rectangle, error) {
	return s.repo.GetAllRectangles()
}

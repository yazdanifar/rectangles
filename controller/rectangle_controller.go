package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rectangles/dto"
	"rectangles/model"
	"rectangles/service"
	"time"
)

type RectangleController struct {
	rectangleService *service.RectangleService
}

func NewRectangleController(rectangleService *service.RectangleService) *RectangleController {
	return &RectangleController{rectangleService: rectangleService}
}

func (c *RectangleController) CreateIntersectingRectangles(ctx *gin.Context) {
	var requestData dto.RectanglesRequest
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON request"})
		return
	}
	main := requestData.Main
	input := requestData.Input
	mainRect := model.Rectangle{X: main.X, Y: main.Y, Width: main.Width, Height: main.Height, Time: time.Now()}
	inputRectangles := make([]model.Rectangle, len(input))
	for i, rect := range input {
		inputRectangles[i] = model.Rectangle{X: rect.X, Y: rect.Y, Width: rect.Width, Height: rect.Height,
			Time: time.Now()}
	}
	intersectingRectangles := c.rectangleService.CreateIntersectingRectangles(mainRect, inputRectangles)
	output := make([]dto.Rectangle, len(intersectingRectangles))
	for i, rect := range intersectingRectangles {
		output[i] = dto.Rectangle{X: rect.X, Y: rect.Y, Width: rect.Width, Height: rect.Height,
			Time: dto.TimeDTO{Time: rect.Time}}
	}
	ctx.JSON(http.StatusOK, output)
}

func (c *RectangleController) GetAllRectangles(ctx *gin.Context) {
	allRectangles, err := c.rectangleService.GetAllRectangles()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Problem while fetching items"})
		return
	}
	output := make([]dto.Rectangle, len(allRectangles))
	for i, rect := range allRectangles {
		output[i] = dto.Rectangle{X: rect.X, Y: rect.Y, Width: rect.Width, Height: rect.Height,
			Time: dto.TimeDTO{Time: rect.Time}}
	}
	ctx.JSON(http.StatusOK, output)
}

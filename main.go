package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"rectangles/controller"
	"rectangles/model"
	"rectangles/repository"
	"rectangles/service"
)

var db *gorm.DB

func main() {
	var err error
	db, err = gorm.Open(sqlite.Open("rectangles.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return
	}

	err = db.AutoMigrate(&model.Rectangle{})
	if err != nil {
		log.Fatalf("Failed to perform migration: %v", err)
		return
	}

	rectangleRepo := repository.NewRectangleRepository(db)
	rectangleService := service.NewRectangleService(rectangleRepo)
	rectangleCtrl := controller.NewRectangleController(rectangleService)

	r := gin.Default()
	r.POST("/", rectangleCtrl.CreateIntersectingRectangles)
	r.GET("/", rectangleCtrl.GetAllRectangles)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}

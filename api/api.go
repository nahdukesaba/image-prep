package api

import (
	v1 "image-prep/api/v1"
	"image-prep/manager"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	imageHandler := v1.NewImageHandler(manager.ImageService())

	rs := gin.Default()
	rs.GET("/convert", imageHandler.ConvertHandler)
	rs.GET("/resize", imageHandler.ResizeHandler)
	rs.GET("/compress", imageHandler.CompressHandler)
	return rs
}

package api

import (
	handlers "github.com/philipphahmann/hack-video-transcoder/internal/adapters/input/api/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, upload *handlers.UploadHandler) {
	r.POST("/api/upload", upload.Upload)
	r.GET("/api/download/:filename", handlers.DownloadHandler)
	r.GET("/api/health", handlers.HealthHandler)
}

package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/philipphahmann/hack-video-transcoder/docs"
	api "github.com/philipphahmann/hack-video-transcoder/internal/adapters/input/api"
	handlers "github.com/philipphahmann/hack-video-transcoder/internal/adapters/input/api/handlers"
	middleware "github.com/philipphahmann/hack-video-transcoder/internal/adapters/input/api/middleware"
	usecases "github.com/philipphahmann/hack-video-transcoder/internal/application/usecases"
	infra "github.com/philipphahmann/hack-video-transcoder/internal/infrastructure"
	"github.com/philipphahmann/hack-video-transcoder/pkg/logger"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	logger.Setup()

	r := gin.New()
	r.Use(middleware.LoggerMiddleware())
	r.Use(gin.Recovery())

	processor := infra.NewFFmpegProcessor()
	useCase := usecases.NewProcessVideoUseCase(processor)
	uploadHandler := handlers.NewUploadHandler(useCase)

	api.RegisterRoutes(r, uploadHandler)

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}

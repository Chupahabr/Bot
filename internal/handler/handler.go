package handler

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"skin-monkey/internal/service"
)

type Handler struct {
	services *service.Service
	log      *slog.Logger
}

func NewHandler(services *service.Service, log *slog.Logger) *Handler {
	return &Handler{
		services: services,
		log:      log,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "https://aim.market/")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400") // Кеширование на 24 часа

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	})

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			skins := v1.Group("/skins")
			{
				skins.POST("/", h.addSkinsHandler)
				skins.GET("/", h.getSkinHandler)
				skins.DELETE("/:id", h.deleteSkinHandler)
			}
		}
	}

	return router
}

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
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		}

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

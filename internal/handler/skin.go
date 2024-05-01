package handler

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"skin-monkey/internal/entity"
	repository "skin-monkey/internal/repository/postgres"
	"time"
)

func (h *Handler) addSkinsHandler(c *gin.Context) {
	h.LogRequest("add skin", c)

	var skin entity.Skin

	if err := c.BindJSON(&skin); err != nil {
		h.log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.services.Skin.CreateSkin(&skin); err != nil {
		h.log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, skin)
}

func (h *Handler) updateSkinHandler(c *gin.Context) {
	h.LogRequest("update skin", c)

	var skin entity.Skin

	if err := c.BindJSON(&skin); err != nil {
		h.log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	skin.Id = c.Param("id")

	if err := h.services.Skin.UpdateSkin(&skin); err != nil {
		h.log.Error(err.Error())
		if err == sql.ErrNoRows {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": repository.ErrSkinNotFound.Error()})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.IndentedJSON(http.StatusOK, skin)
	h.log.Info("update skin", skin)
}

func (h *Handler) deleteSkinHandler(c *gin.Context) {
	h.LogRequest("delete skin", c)

	id := c.Param("id")

	if err := h.services.Skin.DeleteSkin(id); err != nil {
		h.log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "skin deleted"})
	h.log.Info("delete skin",
		slog.String("id", id),
	)
}

func (h *Handler) getSkinHandler(c *gin.Context) {
	h.LogRequest("get skin", c)

	/*var filter entity.SkinFilter
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.log.Info("get skin with filter",
		slog.Any("filter", filter),
	)*/

	skins, err := h.services.Skin.GetSkinFilter()
	if err != nil {
		h.log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, skins)
}

func (h *Handler) LogRequest(message string, c *gin.Context) {
	h.log.Info("Request:"+message,
		slog.String("ip", c.ClientIP()),
		slog.String("time", time.Now().Format("2006-01-02 15:04:05")),
		slog.String("method", c.Request.Method),
	)
}

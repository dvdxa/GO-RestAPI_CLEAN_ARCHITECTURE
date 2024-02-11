package controller

import (
	"net/http"

	"github.com/dvdxa/GO-RestAPI_CLEAN_ARCHITECTURE/config"
	"github.com/dvdxa/GO-RestAPI_CLEAN_ARCHITECTURE/internal/app/pkg/logger"
	"github.com/dvdxa/GO-RestAPI_CLEAN_ARCHITECTURE/internal/app/pkg/models"
	"github.com/dvdxa/GO-RestAPI_CLEAN_ARCHITECTURE/internal/app/usecase"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	cfg     *config.Config
	log     *logger.Logger
	usecase usecase.IService
}

func NewHandler(cfg *config.Config, log *logger.Logger, usecase *usecase.Service) *Handler {
	return &Handler{
		cfg:     cfg,
		log:     log,
		usecase: usecase,
	}
}

func (h *Handler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "app is ok",
	})
}

func (h *Handler) CreateItem(c *gin.Context) {
	var request models.Item

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"err": err.Error(),
		})
		return
	}

	id, err := h.usecase.CreateItem(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) GetItems(c *gin.Context) {
	items, err := h.usecase.GetItems()
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"items": items,
	})
}

func (h *Handler) UpdateItem(c *gin.Context) {
	var request models.Item

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"err": err.Error(),
		})
		return
	}

	if request.ID == nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"err": "id is nil",
		})
		return
	}

	err = h.usecase.UpdateItem(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"msg": "item is updated",
	})
}

package api

import (
	"log"
	"math/rand"
	"net/http"

	"totoconfigapi/internal/db"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	db db.ConfigDB
}

func NewHandler(db db.ConfigDB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	router.GET("/api/configs", CountryLookupMiddleware(), h.GetConfigHandler)
}

func (h *Handler) GetConfigHandler(c *gin.Context) {
	packageID := c.Query("package")
	country_code, _ := c.Get("country")

	country, ok := country_code.(string)
	if !ok {
		country = "ZZ"
	}

	randomPercentile := rand.Intn(100) + 1

	configs, err := h.db.GetConfigs(packageID, country, randomPercentile)
	if err != nil {
		log.Printf("Error during DB access: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve configuration"})
		return
	}

	var response []map[string]interface{}
	for _, config := range configs {
		response = append(response, map[string]interface{}{"main_sku": config.MainSKU})
	}

	c.JSON(http.StatusOK, response)
}

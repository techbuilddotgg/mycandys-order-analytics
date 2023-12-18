package handlers

import (
	"github.com/gin-gonic/gin"
	"mycandys-order-analytics/internal/models"
	"mycandys-order-analytics/internal/services"
)

type Handler struct {
	analytics *services.AnalyticsService
}

func NewHandler() *Handler {
	return &Handler{analytics: services.NewAnalyticsService()}
}

// AddCallEndpoint godoc
// @Summary Add call endpoint
// @Description Add call endpoint
// @Tags Analytics
// @Accept json
// @Produce json
// @Success 200 {object} models.Endpoint
// @Param endpoint body models.Endpoint true "endpoint"
// @Router /analytics [post]
func (h *Handler) AddCallEndpoint(c *gin.Context) {
	var endpoint models.Endpoint

	if err := c.ShouldBindJSON(&endpoint); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.analytics.AddCalledEndpoint(&endpoint); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, endpoint)
}

// GetLastCalledEndpoint godoc
// @Summary Get last called endpoint
// @Description Get last called endpoint
// @Tags Analytics
// @Produce json
// @Success 200 {object} models.Endpoint
// @Router /analytics/last [get]
func (h *Handler) GetLastCalledEndpoint(c *gin.Context) {
	lastCalledEndpoint, err := h.analytics.GetLastCalledEndpoint()

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, lastCalledEndpoint)

}

// GetNumberOfCallForEachEndpoint godoc
// @Summary Get number of call for each endpoint
// @Description Get number of call for each endpoint
// @Tags Analytics
// @Produce json
// @Success 200 {object} map[string]int
// @Router /analytics/each [get]
func (h *Handler) GetNumberOfCallForEachEndpoint(c *gin.Context) {
	numberOfCallForEachEndpoint, err := h.analytics.GetNumberOfCallForEachEndpoint()

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, numberOfCallForEachEndpoint)
}

// GetMostCalledEndpoint godoc
// @Summary Get most called endpoint
// @Description Get most called endpoint
// @Tags Analytics
// @Produce json
// @Success 200
// @Router /analytics/most [get]
func (h *Handler) GetMostCalledEndpoint(c *gin.Context) {
	mostCalledEndpoint, err := h.analytics.GetMostCalledEndpoint()

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, mostCalledEndpoint)
}

// HealthCheck godoc
// @Summary health check
// @Schemes
// @Description do health check
// @Success 200
// @Router /health [get]
// @Tags health
func (h *Handler) HealthCheck(g *gin.Context) {

	health := map[string]string{
		"status": "ok",
	}

	g.JSON(200, health)
}

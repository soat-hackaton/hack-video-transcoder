package handlers

import "github.com/gin-gonic/gin"

// Status godoc
// @Summary Status da API
// @Description Retorna o status do serviço
// @Tags System
// @Produce json
// @Success 200 {object} map[string]interface{} "Health check"
// @Router /api/health [get]
func HealthHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

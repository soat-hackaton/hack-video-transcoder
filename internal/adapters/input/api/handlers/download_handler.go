package handlers

import (
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// DownloadFile godoc
// @Summary Download de arquivo processado
// @Description Faz o download de um arquivo pelo nome
// @Tags Video
// @Produce application/octet-stream
// @Param filename path string true "Nome do arquivo"
// @Success 200 {file} file "Arquivo para download"
// @Failure 404 {object} map[string]string "Arquivo não encontrado"
// @Router /api/download/{filename} [get]
func DownloadHandler(c *gin.Context) {
	filename := c.Param("filename")
	filePath := filepath.Join("outputs", filename)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(404, gin.H{"error": "Arquivo não encontrado"})
		return
	}

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Type", "application/zip")

	c.File(filePath)
}

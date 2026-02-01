package handlers

import (
	"path/filepath"

	"github.com/philipphahmann/hack-video-transcoder/internal/application/ports"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/philipphahmann/hack-video-transcoder/pkg/utils"
)

type UploadHandler struct {
	useCase ports.ProcessVideoUseCase
}

// UploadVideo godoc
// @Summary Upload de vídeo
// @Description Recebe um arquivo de vídeo para processamento
// @Tags Video
// @Accept multipart/form-data
// @Produce json
// @Param video formData file true "Arquivo de vídeo"
// @Success 200 {object} map[string]string "Upload realizado com sucesso"
// @Failure 400 {object} map[string]string "Erro no upload"
// @Router /api/upload [post]
func NewUploadHandler(uc ports.ProcessVideoUseCase) *UploadHandler {
	return &UploadHandler{useCase: uc}
}

func (h *UploadHandler) Upload(c *gin.Context) {
	file, header, err := c.Request.FormFile("video")

	if !utils.IsValidVideoFile(header.Filename) {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Formato de arquivo não suportado",
		})
		return
	}

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	id := uuid.New()
	filename := id.String()
	path := filepath.Join("uploads", filename)

	c.SaveUploadedFile(header, path)

	result := h.useCase.Execute(path, filename)
	c.JSON(200, result)
}

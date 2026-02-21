package handlers

import (
	"log/slog"
	"path/filepath"

	"github.com/philipphahmann/hack-video-transcoder/internal/application/ports"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/philipphahmann/hack-video-transcoder/pkg/logger"
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

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	id := uuid.New()
	filename := id.String() // Usado como task_id

	ctx := logger.WithCorrelationID(c.Request.Context(), filename)
	slog.InfoContext(ctx, "Iniciando processo de solicitação de upload", slog.String("video_filename", header.Filename), slog.String("step", "request_upload_start"))

	if !utils.IsValidVideoFile(header.Filename) {
		slog.ErrorContext(ctx, "Formato de arquivo não suportado", slog.String("video_filename", header.Filename))
		c.JSON(400, gin.H{
			"success": false,
			"message": "Formato de arquivo não suportado",
		})
		return
	}

	path := filepath.Join("uploads", filename)

	err = c.SaveUploadedFile(header, path)
	if err != nil {
		slog.ErrorContext(ctx, "Erro crítico ao salvar arquivo enviado", slog.String("error", err.Error()))
		c.JSON(500, gin.H{"error": "Erro interno ao salvar arquivo"})
		return
	}

	slog.InfoContext(ctx, "Vídeo salvo com sucesso, enviando para processamento", slog.String("path", path))

	result := h.useCase.Execute(ctx, path, filename)

	if result.Success {
		slog.InfoContext(ctx, "Processamento finalizado com sucesso", slog.String("step", "request_upload_success"))
	} else {
		slog.ErrorContext(ctx, "Erro no processamento do vídeo", slog.String("message", result.Message))
	}

	c.JSON(200, result)
}

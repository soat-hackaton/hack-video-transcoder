package video

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"

	video "github.com/philipphahmann/hack-video-transcoder/internal/domain/video"
)

type FFmpegProcessor struct{}

func NewFFmpegProcessor() *FFmpegProcessor {
	return &FFmpegProcessor{}
}

func (p *FFmpegProcessor) Process(ctx context.Context, videoPath, taskID string) video.ProcessingResult {
	slog.InfoContext(ctx, "Iniciando processo de conversão do vídeo", slog.String("videoPath", videoPath), slog.String("task_id", taskID))

	tempDir := filepath.Join("temp", taskID)
	os.MkdirAll(tempDir, 0755)
	defer os.RemoveAll(tempDir)

	framePattern := filepath.Join(tempDir, "frame_%04d.jpg")

	cmd := exec.CommandContext(ctx, "ffmpeg", "-threads", "1", "-i", videoPath, "-vf", "fps=1,scale=-1:720", "-q:v", "2", "-y", framePattern)

	slog.InfoContext(ctx, "Executando comando FFmpeg para extração de frames", slog.String("cmd", cmd.String()))

	outputBytes, err := cmd.CombinedOutput()
	outputStr := string(outputBytes)
	if len(outputStr) > 2000 {
		outputStr = "..." + outputStr[len(outputStr)-2000:]
	}
	if err != nil {

		exitCode := -1
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode = exitErr.ExitCode()
		}

		slog.ErrorContext(ctx, "Falha na execução do FFmpeg", slog.Int("exitCode", exitCode), slog.String("error", err.Error()))

		return video.ProcessingResult{
			Success: false,
			Message: fmt.Sprintf(
				"Erro no ffmpeg (exit=%d): %v\nComando: %s\nOutput:\n%s",
				exitCode,
				err,
				cmd.String(),
				outputStr,
			),
		}
	}

	frames, _ := filepath.Glob(filepath.Join(tempDir, "*.jpg"))
	if len(frames) == 0 {
		slog.ErrorContext(ctx, "Nenhum frame extraído após execução do FFmpeg")
		return video.ProcessingResult{Success: false, Message: "Nenhum frame extraído"}
	}

	slog.InfoContext(ctx, "Frames extraídos com sucesso", slog.Int("frameCount", len(frames)))

	zipPath := filepath.Join("outputs", "frames_"+taskID+".zip")
	slog.InfoContext(ctx, "Iniciando compactação dos frames", slog.String("zipPath", zipPath))

	if err := CreateZip(frames, zipPath); err != nil {
		slog.ErrorContext(ctx, "Erro ao compactar os frames", slog.String("error", err.Error()))
		return video.ProcessingResult{Success: false, Message: err.Error()}
	}

	slog.InfoContext(ctx, "Arquivo ZIP gerado com sucesso", slog.String("zipPath", zipPath))

	return video.ProcessingResult{
		Success:    true,
		Message:    "Processamento concluído",
		ZipPath:    filepath.Base(zipPath),
		FrameCount: len(frames),
	}
}

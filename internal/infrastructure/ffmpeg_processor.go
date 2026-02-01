package video

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	video "github.com/philipphahmann/hack-video-transcoder/internal/domain/video"
)

type FFmpegProcessor struct{}

func NewFFmpegProcessor() *FFmpegProcessor {
	return &FFmpegProcessor{}
}

func (p *FFmpegProcessor) Process(videoPath, timestamp string) video.ProcessingResult {
	tempDir := filepath.Join("temp", timestamp)
	os.MkdirAll(tempDir, 0755)
	defer os.RemoveAll(tempDir)

	framePattern := filepath.Join(tempDir, "frame_%04d.png")

	cmd := exec.Command("ffmpeg", "-i", videoPath, "-vf", "fps=1", "-y", framePattern)
	if output, err := cmd.CombinedOutput(); err != nil {
		return video.ProcessingResult{
			Success: false,
			Message: fmt.Sprintf("Erro no ffmpeg: %s", string(output)),
		}
	}

	frames, _ := filepath.Glob(filepath.Join(tempDir, "*.png"))
	if len(frames) == 0 {
		return video.ProcessingResult{Success: false, Message: "Nenhum frame extraído"}
	}

	zipPath := filepath.Join("outputs", "frames_"+timestamp+".zip")
	if err := CreateZip(frames, zipPath); err != nil {
		return video.ProcessingResult{Success: false, Message: err.Error()}
	}

	return video.ProcessingResult{
		Success:    true,
		Message:    "Processamento concluído",
		ZipPath:    filepath.Base(zipPath),
		FrameCount: len(frames),
	}
}

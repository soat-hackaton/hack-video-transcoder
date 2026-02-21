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

	framePattern := filepath.Join(tempDir, "frame_%04d.jpg")

	cmd := exec.Command("ffmpeg", "-threads", "1", "-i", videoPath, "-vf", "fps=1,scale=-1:720", "-q:v", "2", "-y", framePattern)
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

package ports

import "github.com/philipphahmann/hack-video-transcoder/internal/domain/video"

type ProcessVideoUseCase interface {
	Execute(videoPath, timestamp string) video.ProcessingResult
}

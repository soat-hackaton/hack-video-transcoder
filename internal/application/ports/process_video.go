package ports

import (
	"context"

	"github.com/philipphahmann/hack-video-transcoder/internal/domain/video"
)

type ProcessVideoUseCase interface {
	Execute(ctx context.Context, videoPath, taskID string) video.ProcessingResult
}

package usecases

import (
	"testing"

	"github.com/philipphahmann/hack-video-transcoder/internal/domain/video"
)

type VideoProcessorMock struct {
	Result video.ProcessingResult
}

func (m *VideoProcessorMock) Process(videoPath string, timestamp string) video.ProcessingResult {
	return m.Result
}

func TestProcessVideoUseCase_Success(t *testing.T) {
	mock := &VideoProcessorMock{
		Result: video.ProcessingResult{
			Success: true,
			Message: "ok",
			ZipPath: "file.zip",
		},
	}

	useCase := NewProcessVideoUseCase(mock)

	result := useCase.Execute("video.mp4", "123")

	if !result.Success {
		t.Fatalf("expected success=true")
	}

	if result.ZipPath != "file.zip" {
		t.Fatalf("unexpected zip path")
	}
}

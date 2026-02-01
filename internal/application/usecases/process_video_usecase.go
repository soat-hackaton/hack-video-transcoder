package usecases

import (
	"github.com/philipphahmann/hack-video-transcoder/internal/application/ports"
	video "github.com/philipphahmann/hack-video-transcoder/internal/domain/video"
)

type ProcessVideoUseCaseImpl struct {
	processor video.Processor
}

func NewProcessVideoUseCase(p video.Processor) ports.ProcessVideoUseCase {
	return &ProcessVideoUseCaseImpl{
		processor: p,
	}
}

func (uc *ProcessVideoUseCaseImpl) Execute(videoPath, timestamp string) video.ProcessingResult {
	return uc.processor.Process(videoPath, timestamp)
}

// package usecases

// import (
// 	video "github.com/philipphahmann/hack-video-transcoder/internal/domain/video"
// )

// type Processor interface {
// 	Process(videoPath, timestamp string) video.ProcessingResult
// }

// type ProcessVideoUseCase struct {
// 	processor Processor
// }

// func NewProcessVideoUseCase(p Processor) *ProcessVideoUseCase {
// 	return &ProcessVideoUseCase{processor: p}
// }

// func (uc *ProcessVideoUseCase) Execute(videoPath, timestamp string) video.ProcessingResult {
// 	return uc.processor.Process(videoPath, timestamp)
// }

package video

import "context"

type Processor interface {
	Process(ctx context.Context, videoPath string, timestamp string) ProcessingResult
}

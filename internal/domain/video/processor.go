package video

type Processor interface {
	Process(videoPath string, timestamp string) ProcessingResult
}

package handlers

import (
	"bytes"
	"context"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/philipphahmann/hack-video-transcoder/internal/domain/video"
)

type ProcessVideoUseCaseMock struct{}

func (m *ProcessVideoUseCaseMock) Execute(ctx context.Context, videoPath, timestamp string) video.ProcessingResult {
	return video.ProcessingResult{
		Success: true,
		Message: "ok",
	}
}

func TestUploadHandler_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, _ := writer.CreateFormFile("video", "test.mp4")
	part.Write([]byte("fake video content"))

	writer.Close()

	req, _ := http.NewRequest(http.MethodPost, "/api/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp := httptest.NewRecorder()

	router := gin.New()
	handler := NewUploadHandler(&ProcessVideoUseCaseMock{})
	router.POST("/api/upload", handler.Upload)

	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.Code)
	}
}

func TestUploadHandler_InvalidFile(t *testing.T) {
	gin.SetMode(gin.TestMode)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, _ := writer.CreateFormFile("video", "test.txt")
	part.Write([]byte("fake video content"))

	writer.Close()

	req, _ := http.NewRequest(http.MethodPost, "/api/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp := httptest.NewRecorder()

	router := gin.New()
	handler := NewUploadHandler(&ProcessVideoUseCaseMock{})
	router.POST("/api/upload", handler.Upload)

	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", resp.Code)
	}
}

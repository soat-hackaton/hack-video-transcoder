package utils

import "testing"

func TestIsValidVideoFile(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		expected bool
	}{
		{"MP4 válido", "video.mp4", true},
		{"AVI válido", "video.avi", true},
		{"Maiúsculo", "VIDEO.MKV", true},
		{"Arquivo inválido", "document.pdf", false},
		{"Sem extensão", "video", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidVideoFile(tt.filename)
			if result != tt.expected {
				t.Errorf("esperado %v, recebido %v", tt.expected, result)
			}
		})
	}
}

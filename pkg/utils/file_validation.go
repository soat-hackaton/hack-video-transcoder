package utils

import (
	"path/filepath"
	"strings"
)

// IsValidVideoFile valida se o arquivo possui uma extensão de vídeo suportada
func IsValidVideoFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))

	validExts := map[string]bool{
		".mp4":  true,
		".avi":  true,
		".mov":  true,
		".mkv":  true,
		".wmv":  true,
		".flv":  true,
		".webm": true,
	}

	return validExts[ext]
}

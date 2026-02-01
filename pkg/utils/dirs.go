package utils

import "os"

// CreateRequiredDirs cria os diretórios necessários para a aplicação
func CreateRequiredDirs() error {
	dirs := []string{
		"uploads",
		"outputs",
		"temp",
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	return nil
}

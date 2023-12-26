package inputs

import (
	"log"

	f "github.com/Hakitsyu/logs-processor/internal/file"
)

const (
	FileInputType = "file"
)

type FileInput struct {
	config *FileInputConfig
}

type FileInputConfig struct {
	Paths []string
}

func (fileInput *FileInput) getPaths() []string {
	return fileInput.config.Paths
}

func (fileInput *FileInput) Load() [][][]byte {
	paths := fileInput.getPaths()
	result := make([][][]byte, len(paths))

	for _, path := range paths {
		lines, err := fileInput.load(path)
		if err != nil {
			log.Fatal(err)
		}

		result = append(result, lines)
	}

	return result
}

func (fileInput *FileInput) load(path string) ([][]byte, error) {
	return f.ReadLines(path)
}

func NewFileInput(config *FileInputConfig) *FileInput {
	return &FileInput{
		config: config,
	}
}

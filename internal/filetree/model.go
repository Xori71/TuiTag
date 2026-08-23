package filetree

import (
	"os"
)

type Model struct {
	currentDir string
	fileList   []os.DirEntry
	index      int
}

func initModel() (Model, error) {
	// Open current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		return Model{}, err
	}

	entries, err := os.ReadDir(currentDir)

	return Model{
		currentDir: currentDir,
		fileList:   entries,
		index:      0,
	}, err
}

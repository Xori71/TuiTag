package filetree

import (
	"os"
)

type Model struct {
	currentDir string
	fileList   []os.DirEntry
	index      int
}

func New(savedDir string) (Model, error) {
	var currentDir string
	var err error
	if savedDir == "" {
		currentDir, err = os.Getwd()
	} else {
		currentDir = savedDir
	}
	if err != nil {
		return Model{}, err
	}

	entries, err := os.ReadDir(currentDir)
	if err != nil {
		return Model{}, nil
	}

	return Model{
		currentDir: currentDir,
		fileList:   entries,
		index:      0,
	}, nil
}

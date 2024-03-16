package images

import (
	"io"
	"io/fs"
	"os"
)

func ImageToByteArray(imagePath string) ([]byte, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func GetAllImagesAsBytes(fileDir string) ([][]float64, error) {
	filesInDir, err := os.ReadDir(fileDir)
	if err != nil {
		return nil, err
	}

	files, err := getFileArray(filesInDir, fileDir)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func getFileArray(filesInDir []fs.DirEntry, fileDir string) ([][]float64, error) {
	images := make([][]float64, len(filesInDir))
	for i := range filesInDir {
		if filesInDir[i].IsDir() {
			continue
		}
		imageBytes, err := ImageToByteArray(fileDir + filesInDir[i].Name())
		if err != nil {
			return nil, err
		}

		inputs, err := ImageStringToArray(imageBytes)
		if err != nil {
			return nil, err
		}
		images[i] = inputs
	}
	return images, nil
}

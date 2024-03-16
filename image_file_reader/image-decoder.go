package image_file_reader

import (
	"bytes"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	. "simple-neural-network/utils"
)

func ImageStringToArray(byteArray []byte) ([]float64, error) {
	pngImage, err := decodePNG(byteArray)
	if err != nil {
		return nil, err
	}

	twoDimensionalArray := imageTo2DArray(pngImage)
	return Flatten(twoDimensionalArray), nil
}

// This function assumes `imageBytes` is a slice of bytes containing the PNG data.
func decodePNG(imageBytes []byte) (image.Image, error) {
	img, _, err := image.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		return nil, err
	}
	return img, nil
}

func imageTo2DArray(img image.Image) [][]float64 {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	result := make([][]float64, height)
	for y := 0; y < height; y++ {
		result[y] = make([]float64, width)
		for x := 0; x < width; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			// Convert color to grayscale using the luminance method
			// Note: r, g, b are uint32 but within a 0-65535 range,
			// which corresponds to their 0-255 value but multiplied by 256 (or left-shifted by 8 bits).
			// Dividing by 256 (or right-shifting by 8 bits) normalizes them back to 0-255 before calculating gray.
			gray := 0.299*float64(r>>8) + 0.587*float64(g>>8) + 0.114*float64(b>>8)
			// Normalize to range 0.0 - 1.0. Now gray is in the range of 0-255, so divide by 255.
			result[y][x] = gray / 255.0
		}
	}

	return result
}

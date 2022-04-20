package main

import (
	"image"
	"image/png"
	"log"
	"math"
	"os"
	"place/data"
)

func main() {
	w, h, pixels, err := imgToPixels("./in.png")
	if err != nil {
		log.Fatal(err)
	}

	err = makeNewImage(w, h, pixels)

	if err != nil {
		log.Fatal(err)
	}
}

func makeNewImage(width int, height int, pixels []data.Pixel) error {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for _, pixel := range pixels {
		img.Set(pixel.X, pixel.Y, data.RgbToRgba(pixel.Color.RGB))
	}

	out, err := os.Create("out.png")
	if err != nil {
		return err
	}

	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(out)

	err = png.Encode(out, img)
	if err != nil {
		return err
	}
	return nil
}

func imgToPixels(path string) (int, int, []data.Pixel, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, 0, nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	imgData, _, err := image.Decode(file)
	if err != nil {
		return 0, 0, nil, err
	}

	bounds := imgData.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var pixels []data.Pixel

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			color, err := approxColor(data.RgbaToRgb(imgData.At(x, y).RGBA()))
			if err != nil {
				return 0, 0, nil, err
			}

			if color.Name != data.Transparent.Name {
				pixels = append(pixels, data.Pixel{
					Color: color,
					X:     x,
					Y:     y,
				})
			}

		}
	}

	return width, height, pixels, nil
}

func approxColor(r int, g int, b int, a int) (data.Color, error) {

	if a == 0 {
		return data.Transparent, nil
	}

	rgb := []int{r, g, b}

	distances := map[string]float64{}

	for _, color := range data.Colors() {
		distances[color.Name] = l2(rgb, color.RGB)
	}

	var closest data.Color
	minVal := math.Pow(2, 30)
	for k, v := range distances {
		if v < minVal {

			color, err := data.StringToColor(k)
			if err != nil {
				return data.Color{}, err
			}
			closest = color
			minVal = v
		}
	}
	return closest, nil
}

func l2(listA []int, listB []int) float64 {
	return math.Sqrt(math.Pow(float64(listA[0]-listB[0]), 2) +
		math.Pow(float64(listA[1]-listB[1]), 2) +
		math.Pow(float64(listA[2]-listB[2]), 2))
}

package neural

import (
	"encoding/csv"
	"fmt"
	"image/png"
	"io"
	"os"
	"strconv"
	"testing"
)

func TestDigit(t *testing.T) {
	model := Model{
		Config: ModelConfig{
			LayerCount:   4,
			NodeCount:    []int{28 * 28, 200, 40, 10},
			LearningRate: 0.1,
			Activation:   Sigmoid,
		},
	}
	model.Init()

	file, err := os.Open("digit-data/train.csv")
	panicErr(err)
	defer file.Close()
	reader := csv.NewReader(file)
	_, err = reader.Read() // header
	panicErr(err)
	count := 0
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		panicErr(err)

		imgFile := fmt.Sprintf("digit-data/Images/train/%s", row[0])
		label, err := strconv.Atoi(row[1])
		panicErr(err)

		imgReader, err := os.Open(imgFile)
		panicErr(err)

		img, err := png.Decode(imgReader)
		panicErr(err)

		input := make([]float64, 28*28)
		for i := 0; i < 28; i++ {
			for j := 0; j < 28; j++ {
				color := img.At(i, j)
				r, g, b, _ := color.RGBA()
				if r+g+b > 255 {
					input[i*28+j] = 0
				} else {
					input[i*28+j] = 1
				}
			}
		}
		output := make([]float64, 10)
		output[label] = 1
		for i := 0; i < 5; i++ {
			model.Feed(input, output)
		}
		fmt.Printf("%5d Error %.5f\n", count, model.Value.Error)
		count++
	}
}

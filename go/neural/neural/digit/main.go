package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/xdean/share/go/neural/neural"
	"image/png"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	Test()
}

func Test() {
	model := neural.Model{
		Config: neural.ModelConfig{
			LayerCount:   4,
			NodeCount:    []int{28 * 28, 200, 40, 10},
			LearningRate: 0.1,
			Activation:   neural.Sigmoid,
		},
	}
	err := os.Chdir("go/neural/neural/digit")
	neural.PanicErr(err)
	err = model.Load("output/model/digit-10000.model")
	neural.PanicErr(err)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter image number (49000-69999) or exit")
		line, err := reader.ReadString('\n')
		if err == io.EOF || strings.Contains(line, "exit") {
			fmt.Println("Bye")
			break
		}
		line = strings.Trim(line, "\n \t")
		imgFile := fmt.Sprintf("data/Images/test/%s.png", line)
		input := DigitReadImage(imgFile)
		predict := model.Predict(input)
		fmt.Println("Predict:", predict)
		max := 0
		for i, v := range predict {
			if v > predict[max] {
				max = i
			}
		}
		fmt.Printf("Answer: %d\n", max)
	}
}

func Train() {
	model := neural.Model{
		Config: neural.ModelConfig{
			LayerCount:   4,
			NodeCount:    []int{28 * 28, 200, 40, 10},
			LearningRate: 0.1,
			Activation:   neural.Sigmoid,
		},
	}
	model.Init()

	file, err := os.Open("data/train.csv")
	neural.PanicErr(err)
	defer file.Close()
	reader := csv.NewReader(file)
	_, err = reader.Read() // header
	neural.PanicErr(err)
	count := 0
	for {
		row, err := reader.Read()
		if err == io.EOF || count > 10000 {
			break
		}
		neural.PanicErr(err)

		imgFile := fmt.Sprintf("data/Images/train/%s", row[0])
		label, err := strconv.Atoi(row[1])
		neural.PanicErr(err)

		input := DigitReadImage(imgFile)
		output := make([]float64, 10)
		output[label] = 1
		for i := 0; i < 5; i++ {
			model.Feed(input, output)
		}
		fmt.Printf("%5d Error %.5f\n", count, model.Value.Error)
		count++
	}

	err = model.Save("output/model/digit-10000.model")
	neural.PanicErr(err)
}

func DigitReadImage(imgFile string) []float64 {
	imgReader, err := os.Open(imgFile)
	neural.PanicErr(err)
	img, err := png.Decode(imgReader)
	neural.PanicErr(err)
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
	return input
}

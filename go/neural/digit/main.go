package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/xdean/goex/xgo"
	"github.com/xdean/share/go/neural/neural"
	"image/png"
	"io"
	"os"
	"strconv"
	"strings"
)

var model = neural.Model{
	Config: neural.ModelConfig{
		LayerCount:   4,
		NodeCount:    []int{28 * 28, 200, 40, 10},
		LearningRate: 0.1,
		Activation:   neural.Sigmoid,
	},
}
var modelPath = "digit-all-200x40.model"

func main() {
	wd, err := os.Getwd()
	xgo.MustNoError(err)
	if strings.HasSuffix(wd, "go/neural/digit") {
		err := os.Chdir("go/neural/digit")
		xgo.MustNoError(err)
	}

	train := flag.Bool("train", false, "Train model")
	flag.StringVar(&modelPath, "model", modelPath, "Model path")
	flag.Parse()

	if *train {
		Train()
	} else {
		Test()
	}
}

func Test() {
	err := model.Load(modelPath)
	xgo.MustNoError(err)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter image number (49000-69999) or exit")
		line, err := reader.ReadString('\n')
		if err == io.EOF || strings.Contains(line, "exit") {
			fmt.Println("Bye")
			break
		}
		line = strings.Trim(line, "\n\r \t")
		imgFile := fmt.Sprintf("data/Images/test/%s.png", line)
		fmt.Println(imgFile)
		input, err := DigitReadImage(imgFile)
		if err != nil {
			fmt.Println("Fail to read image:", err)
			continue
		}
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
	model.Init()

	file, err := os.Open("data/train.csv")
	xgo.MustNoError(err)
	defer file.Close()
	reader := csv.NewReader(file)
	_, err = reader.Read() // header
	xgo.MustNoError(err)
	count := 0
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		xgo.MustNoError(err)

		imgFile := fmt.Sprintf("data/Images/train/%s", row[0])
		label, err := strconv.Atoi(row[1])
		xgo.MustNoError(err)

		input, err := DigitReadImage(imgFile)
		xgo.MustNoError(err)
		output := make([]float64, 10)
		output[label] = 1
		for i := 0; i < 5; i++ {
			model.Feed(input, output)
		}
		fmt.Printf("%5d Error %.5f\n", count, model.Value.Error)
		count++
	}

	err = model.Save(modelPath)
	xgo.MustNoError(err)
}

func DigitReadImage(imgFile string) ([]float64, error) {
	imgReader, err := os.Open(imgFile)
	if err != nil {
		return nil, err
	}
	img, err := png.Decode(imgReader)
	if err != nil {
		return nil, err
	}
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
	return input, nil
}

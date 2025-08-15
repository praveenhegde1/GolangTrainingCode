package main

import (
	"fmt"
	"log"
	"sync"

	"gonum.org/v1/gonum/mat"
)

type StockData struct {
	Prices  []float64
	Volumes []int
}

func LoadData(filename string) (*StockData, error) {
	/*file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	data := &StockData{}
	for _, row := range rows {
		price, err := strconv.ParseFloat(row[1], 64)
		if err != nil {
			return nil, err
		}
		data.Prices = append(data.Prices, price)

		volume, err := strconv.Atoi(row[2])
		if err != nil {
			return nil, err
		}
		data.Volumes = append(data.Volumes, volume)
	}

	return data, nil */

	return &StockData{Prices: []float64{50.0, 60.0, 70.0, 80.0, 90.0}, Volumes: []int{100, 200, 300, 400, 500}}, nil
}

func MakePrediction(model *mat.VecDense, input []float64) (float64, error) {

	return 55.0, nil
}

func TrainModel(data *StockData) (*mat.VecDense, error) {
	ST := mat.NewDense(len(data.Prices), 2, nil)
	for i := range data.Prices {
		ST.Set(i, 0, data.Prices[i])
		ST.Set(i, 1, float64(data.Volumes[i]))
	}
	return mat.NewVecDense(len(data.Prices), nil), nil
}

func main() {

	data, err := LoadData("stack.csv")
	if err != nil {
		log.Fatal(err)
	}

	model, err := TrainModel(data)

	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	for i := 0; i <= 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			input := []float64{float64(i), float64(i*100)}
			prediction, err := MakePrediction(model, input)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Prediction for day %d: %f\n", i, prediction)

		}(i)
		wg.Wait()
	}
}

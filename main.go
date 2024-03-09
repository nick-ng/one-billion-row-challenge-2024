package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type weatherStation struct {
	Hash  int
	City  string
	Max   float64
	Min   float64
	Mean  float64
	Count int64
}

func main() {
	start := time.Now()
	// load file
	file, err := os.Open("./data/measurements.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// setup reader
	reader := bufio.NewReader(file)

	// results array
	var myWeatherStations = [10500]weatherStation{}

	// loop
	keepGoing := true
	linesProcessed := 0

	for keepGoing {
		tempBytes, _, err := reader.ReadLine()
		tempLine := string(tempBytes)

		if err != nil {
			fmt.Println("error:", err)
			keepGoing = false
			continue
		}

		temp := strings.Split(tempLine, ";")

		if len(temp) != 2 {
			fmt.Println("invalid line", tempLine)
			os.Exit(1)
		}

		city := temp[0]
		temperatureString := temp[1]
		temperature, err := strconv.ParseFloat(temperatureString, 64)

		if err != nil {
			fmt.Println("invalid temperature", temperatureString)
		}

		tempWeatherStation := getWeatherStation(city, &myWeatherStations)

		tempWeatherStation.Max = math.Max(tempWeatherStation.Max, temperature)
		tempWeatherStation.Min = math.Min(tempWeatherStation.Min, temperature)
		tempWeatherStation.Mean = (tempWeatherStation.Mean*float64(tempWeatherStation.Count) + temperature) / float64(tempWeatherStation.Count+1)
		tempWeatherStation.Count = tempWeatherStation.Count + 1

		setWeatherStation(tempWeatherStation, &myWeatherStations)
		linesProcessed++
		if linesProcessed%100000 == 0 {
			fmt.Printf("Progress: %0.2f%%\n", float64(linesProcessed)/float64(1_000_000_000)*100)
		}
	}

	weatherStationsSlice := []weatherStation{}

	for i := 0; i < len(myWeatherStations); i++ {
		myWeatherStation := myWeatherStations[i]
		if len(myWeatherStation.City) > 0 {
			weatherStationsSlice = append(weatherStationsSlice, myWeatherStation)

			// if i != myWeatherStation.Hash {
			// 	fmt.Println(i, myWeatherStations[i])
			// }
		}
	}

	sort.Slice(weatherStationsSlice, func(i int, j int) bool {
		a := weatherStationsSlice[i]
		b := weatherStationsSlice[j]

		return a.City < b.City
	})

	for i := 0; i < len(weatherStationsSlice); i++ {
		weatherStation := weatherStationsSlice[i]

		fmt.Printf("%s=%0.1f/%0.1f/%0.1f\n", weatherStation.City, weatherStation.Min, weatherStation.Mean, weatherStation.Max)

	}

	elapsed := time.Since(start)
	fmt.Printf("%d lines took %s\n", linesProcessed, elapsed)
}

func getWeatherStation(city string, weatherStations *[10500]weatherStation) weatherStation {
	initialPosition := hashCityName(city)

	for i := 0; i < 10500; i++ {
		actualPosition := (initialPosition + i) % 10500

		tempWeatherStation := weatherStations[actualPosition]

		if tempWeatherStation.City == city {
			return tempWeatherStation
		}
	}

	return weatherStation{
		Hash:  initialPosition,
		City:  city,
		Max:   -math.MaxFloat64,
		Min:   math.MaxFloat64,
		Mean:  0,
		Count: 0,
	}
}

func setWeatherStation(newWeatherStation weatherStation, weatherStations *[10500]weatherStation) {
	city := newWeatherStation.City

	initialPosition := hashCityName(city)

	for i := 0; i < 10500; i++ {
		actualPosition := (initialPosition + i) % 10500

		tempWeatherStation := weatherStations[actualPosition]

		if tempWeatherStation.City == city || len(tempWeatherStation.City) == 0 {
			weatherStations[actualPosition] = newWeatherStation
			return
		}
	}

	fmt.Println("problem setting weather station")
	os.Exit(1)
}

func hashCityName(city string) int {
	temp := 0

	for i := 0; i < len(city); i++ {
		a := int(city[i])
		temp = temp + (a * i)
	}

	return temp % 10500
}

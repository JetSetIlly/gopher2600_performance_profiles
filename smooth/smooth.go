package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func readDataFromFile(filename string) ([]float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []float64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue // Skip comments
		}
		fields := strings.Fields(line)
		if len(fields) < 3 {
			continue // Skip malformed lines
		}
		value, err := strconv.ParseFloat(fields[2], 64)
		if err != nil {
			return nil, err
		}
		data = append(data, value)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func absoluteDifferences(data []float64) []float64 {
	differences := make([]float64, 0, len(data)-1)
	for i := 1; i < len(data); i++ {
		differences = append(differences, math.Abs(data[i]-data[i-1]))
	}
	return differences
}

func variance(data []float64) float64 {
	mean := 0.0
	for _, value := range data {
		mean += value
	}
	mean /= float64(len(data))

	variance := 0.0
	for _, value := range data {
		variance += math.Pow(value-mean, 2)
	}
	variance /= float64(len(data))

	return variance
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run smoothness.go file1 file2")
		os.Exit(1)
	}

	file1 := os.Args[1]
	file2 := os.Args[2]

	data1, err := readDataFromFile(file1)
	if err != nil {
		log.Fatalf("Error reading data from %s: %v", file1, err)
	}

	data2, err := readDataFromFile(file2)
	if err != nil {
		log.Fatalf("Error reading data from %s: %v", file2, err)
	}

	diff1 := absoluteDifferences(data1)
	diff2 := absoluteDifferences(data2)

	variance1 := variance(diff1)
	variance2 := variance(diff2)

	fmt.Printf("Variance of differences for %s: %.6f\n", file1, variance1)
	fmt.Printf("Variance of differences for %s: %.6f\n", file2, variance2)

	if variance1 < variance2 {
		fmt.Println("File", file1, "is smoother")
	} else if variance1 > variance2 {
		fmt.Println("File", file2, "is smoother")
	} else {
		fmt.Println("Equal smoothness")
	}
}

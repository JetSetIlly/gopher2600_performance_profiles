package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
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

	clean := func(s string) string {
		s = filepath.Base(s)
		s, _ = strings.CutSuffix(s, filepath.Ext(s))
		return s
	}

	display1 := clean(file1)
	display2 := clean(file2)

	diff1 := absoluteDifferences(data1)
	diff2 := absoluteDifferences(data2)

	variance1 := variance(diff1)
	variance2 := variance(diff2)

	fmt.Println("Variance of Differences")
	fmt.Printf("\t%.6f: %s\n", variance1, display1)
	fmt.Printf("\t%.6f: %s\n", variance2, display2)

	if variance1 < variance2 {
		fmt.Println(display1, "is smoother")
	} else if variance1 > variance2 {
		fmt.Println(display2, "is smoother")
	} else {
		fmt.Println("Both sets of data have equal smoothness")
	}
}

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var data []float64
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			continue
		}
		data = append(data, number)
		if len(data) == 1 {
			continue
		}
		start := len(data) - 4
		if start > 0 {
			data = data[start:]
		}
		size := float64(len(data))
		Sx, Sy, Sxy, P_Sx, P_Sy := CalcSums(data)
		PCC := PearsonCorrelationCoefficient(size, Sx, Sy, Sxy, P_Sx, P_Sy)

		if PCC > 1 {
			PCC = 1
		} else if PCC < -1 {
			PCC = -1
		}

		b := LinearStat(data, size, Sx, Sy, Sxy, P_Sx)

		min, max := Guess(b, PCC)
		fmt.Println(min, max)
	}
}

func CalcSums(data []float64) (float64, float64, float64, float64, float64) {
	Sx := float64(0)
	Sy := float64(0)
	Sxy := float64(0)
	P_Sx := float64(0)
	P_Sy := float64(0)
	for i := 0; i < len(data); i++ {
		Sx += float64(i)
		Sy += data[i]
		Sxy += float64(i) * data[i]
		P_Sx += float64(math.Pow(float64(i), 2))
		P_Sy += math.Pow(data[i], 2)
	}
	return Sx, Sy, Sxy, P_Sx, P_Sy
}

func LinearStat(data []float64, size float64, nbrs ...float64) float64 {
	numerator := (size * nbrs[2]) - (nbrs[0] * nbrs[1])
	denominator := (size * nbrs[3]) - math.Pow(nbrs[0], 2)
	slope := numerator / denominator
	b := (nbrs[1] - (slope * nbrs[0])) / size
	return b
}

func PearsonCorrelationCoefficient(size float64, n ...float64) float64 {
	numerator := (size * n[2]) - (n[0] * n[1])
	denominator := ((size * n[3]) - (n[0] * n[0])) * ((size * n[4]) - (n[1] * n[1]))
	if denominator == 0 {
		return 0
	}
	return numerator / math.Sqrt(denominator)
}

func Guess(mx, PCC float64) (float64, float64) {
	max := mx + (PCC * 25) + 20
	min := max - 40

	if min > max {
		min, max = max, min
	}

	return min, max
}

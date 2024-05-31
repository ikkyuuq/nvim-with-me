package math

func Average(i []float64) float64 {
	total := float64(0)
	for _, x := range i {
		total += x
	}
	return total / float64(len(i))
}

func Min(i []float64) float64 {
	min := i[0]
	for _, x := range i {
		if x < min {
			min = x
		}
	}
	return min
}

func Max(i []float64) float64 {
	max := i[0]
	for _, x := range i {
		if x > max {
			max = x
		}
	}
	return max
}

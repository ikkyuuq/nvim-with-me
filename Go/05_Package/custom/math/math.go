package math

func Average(i []float64) float64 {
  total := float64(0)
  for _, x := range i {
    total += x
  }
  return total / float64(len(i))
}

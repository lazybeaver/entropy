// Compute Shannon Entropy of a byte stream
// H = - Σ P(x) * log P(x)

package entropy

import (
	"io"
	"math"
	"strings"
)

type shannon struct {
	frequencies map[byte]int
	total       int
}

func (s *shannon) Init() {
	s.frequencies = make(map[byte]int)
	for i := 0; i < 256; i++ {
		s.frequencies[byte(i)] = 0
	}
	s.total = 0
}

func (s *shannon) Write(data []byte) (int, error) {
	var count int
	for _, b := range data {
		s.frequencies[b] += 1
		count += 1
	}
	s.total += count
	return count, nil
}

func (s *shannon) Value() float64 {
	var entropy float64
	for _, count := range s.frequencies {
		if count > 0 {
			pval := float64(count) / float64(s.total)
			pinv := float64(s.total) / float64(count)
			entropy += pval * math.Log2(pinv)
		}
	}
	return entropy
}

func NewShannonEstimator() Estimator {
	s := &shannon{}
	s.Init()
	return s
}

func Shannon(s string) (float64, error) {
	reader := strings.NewReader(s)
	estimator := NewShannonEstimator()
	_, err := io.Copy(estimator, reader)
	if err != nil {
		return 0.0, err
	}
	return estimator.Value(), nil
}

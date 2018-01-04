package request

import (
	"strconv"
	"time"
	"math/rand"
)

type Math struct {
	n float64
}

func (m *Math) Random() *Math {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	m.n = r.Float64()
	return m
}

func (m *Math) ToFixed(n int) float64 {
	number := strconv.FormatFloat(m.n, 'f', n, 64)
	m.n, _ = strconv.ParseFloat(number, 64)
	return m.n
}

func (m *Math) Multiply(n float64) *Math {
	m.n *= n
	return m
}

func (m *Math) Plus(n float64) *Math {
	m.n += n
	return m
}

func (m *Math) Minus(n float64) *Math {
	m.n = n - m.n
	return m
}


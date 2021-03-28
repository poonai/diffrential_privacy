package differential_privacy

import (
	"math"
)

type CMSServer struct {
	// m is the domain length.
	m int
	c float64
	// kWiseHash is the list of hash function used to map the data on the domain length m.
	kWiseHash []HashMapper
	// matrix is the sketch matrix which holds the aggregation of client donated sampled data.
	matrix [][]float64
	// collected samples.
	n int
}

func GenMatric(i, j int) [][]float64 {
	mat := make([][]float64, i)
	for x := 0; x < i; x++ {
		mat[x] = make([]float64, j)
	}
	return mat
}
func calculateC(epsilon float64) float64 {
	return ((math.Pow(math.E, epsilon/2) + 1) / (math.Pow(math.E, epsilon/2) - 1))
}

func (s *CMSServer) Track(ev *EncodedValue) {
	for i, x := range ev.V {
		s.matrix[ev.J][i] += (float64(len(s.kWiseHash)) * (((s.c / 2) * x) + (0.5 * 1)))
	}
	s.n += 1
}

func (s *CMSServer) Estimate(in []byte) float64 {
	freq := float64(0)
	for i, v := range s.matrix {
		freq += v[s.kWiseHash[i].Calculate(in)]
	}
	return (float64(s.m) / (float64(s.m) - 1) * (((1 / float64(len(s.kWiseHash))) * freq) - (float64(s.n) / float64(s.m))))
}

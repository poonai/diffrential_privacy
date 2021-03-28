package differential_privacy

import (
	"math"
	"math/rand"

	"gonum.org/v1/gonum/stat/distuv"
)

// HashMapper Calculate Hash for the given data and plot the hash for the
// determined domain length m.
type HashMapper interface {
	Calculate(data []byte) int
}

// calculateProb return the probability. Higher the epsilon higher the probability.
func calculateProb(epsilon int) float64 {
	return float64(1 / (1 + math.Pow(math.E, float64(epsilon)/2)))
}

type CMSClient struct {
	// m is the domain length.
	m int
	// kWiseHash is the list of hash function used to map the data on the domain length m.
	kWiseHash []HashMapper
	// prob is used
	prob float64
}

type EncodedValue struct {
	V []float64
	J int
}

// Encode converts the user data into vector v. which is later used by server to construct the
// cms matrix to calculate the frequency of the all the data d in the domain D.
func (c *CMSClient) Encode(data []byte) *EncodedValue {
	v := c.initializeV()
	j := rand.Intn(len(c.kWiseHash))
	// flip the -1 to 1 on the sampled j on the initialized vector.
	v[c.kWiseHash[j].Calculate(data)] = 1
	for i := range v {
		if distuv.UnitUniform.Rand() < float64(c.prob) {
			// flip if the incoming sample is less than the probability.
			v[i] *= -1
		}
	}
	return &EncodedValue{
		V: v,
		J: j,
	}
}

func (c *CMSClient) initializeV() []float64 {
	v := make([]float64, 0, c.m)
	for i := 0; i < c.m; i++ {
		v = append(v, -1)
	}
	return v
}

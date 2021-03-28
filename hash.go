package differential_privacy

import "github.com/OneOfOne/xxhash"

type hasher struct {
	// m is the domain length.
	m uint32

	xxhasher *xxhash.XXHash32
}

func (h *hasher) Calculate(data []byte) int {
	h.xxhasher.Reset()
	h.xxhasher.Write(data)
	return int(h.xxhasher.Sum32() % h.m)
}

func GenHash(j int, m uint32) []HashMapper {
	h := make([]HashMapper, 0, j)
	for i := 0; i < j; i++ {
		h = append(h, &hasher{
			m:        m,
			xxhasher: xxhash.NewS32(uint32(i)),
		})
	}
	return h
}

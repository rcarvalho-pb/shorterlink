package bloomfilter

import (
	"hash"
	"hash/fnv"
	"math"
)

type BloomFilter struct {
	Counts    []uint8
	HashFuncs []hash.Hash64
	Size      uint
	Inserted  uint
	Capacity  uint
	HashSeeds []uint32
}

func NewBloomFilter(capacity uint, accuracy float64) *BloomFilter {
	optimalSize := optimalSize(capacity, accuracy)
	numHashes := optimalHashes(optimalSize, capacity)
	hashFuncs := make([]hash.Hash64, numHashes)
	for i := range numHashes {
		hashFuncs[i] = fnv.New64()
	}
	return &BloomFilter{
		Counts:    make([]uint8, optimalSize),
		HashFuncs: hashFuncs,
		Size:      optimalSize,
		Inserted:  0,
		Capacity:  capacity,
		HashSeeds: getSeeds(len(hashFuncs)),
	}
}

func (bf *BloomFilter) hashItem(item string, seed int) uint64 {
	h := bf.HashFuncs[seed]
	h.Reset()
	h.Write([]byte{byte(bf.HashSeeds[seed])})
	h.Write([]byte(item))
	return h.Sum64()
}

func (bf *BloomFilter) Add(data string) {
	if bf.Inserted >= bf.Capacity/2 {
		bf.resize()
	}
	for i := range len(bf.HashFuncs) {
		position := bf.hashItem(data, i) % uint32(bf.Size)
	}
}

func getSeeds(size int) []uint32 {
	seeds := make([]uint32, size)
	seeds[0] = 1
	seeds[1] = 1
	for i := 2; i < size; i++ {
		seeds[i] = seeds[i-1] + seeds[i-2]
	}
	return seeds
}

func optimalSize(n uint, p float64) uint {
	return uint(math.Ceil(-1 * float64(n) * math.Log(p) / (math.Pow(math.Log(2), 2))))
}

func optimalHashes(m uint, n uint) int {
	return int(math.Ceil(float64(m) / float64(n) * math.Log(2)))
}

package utils

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func GenerateRandomBrowserSize(minSize, maxSize int) (width, height int64) {
	aspectRatios := []float64{16.0 / 9.0, 16.0 / 10.0, 4.0 / 3.0}

	rand.Shuffle(len(aspectRatios), func(i, j int) {
		aspectRatios[i], aspectRatios[j] = aspectRatios[j], aspectRatios[i]
	})

	aspectRatio := aspectRatios[0]

	width = int64(rand.Intn(maxSize-minSize+1) + minSize)
	height = int64(float64(width) / aspectRatio)

	return width, height
}

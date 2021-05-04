package vector_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/tmhmitchell/ebitoolbox/datastructures/vector"
)

func TestEuclidianDistance(t *testing.T) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	v0 := vector.NewVec2(rng.NormFloat64()*100, rng.NormFloat64()*100)
	v1 := vector.NewVec2(rng.NormFloat64()*100, rng.NormFloat64()*100)

	dist0to1 := v0.EuclidianDistance(v1)
	dist1to0 := v1.EuclidianDistance(v0)

	if dist0to1 != dist1to0 {
		t.Errorf("Distances are not the same: %f, %f", dist0to1, dist1to0)
	}
}

func TestManhattanDistance(t *testing.T) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	v0 := vector.NewVec2(rng.NormFloat64()*100, rng.NormFloat64()*100)
	v1 := vector.NewVec2(rng.NormFloat64()*100, rng.NormFloat64()*100)

	dist0to1 := v0.ManhattanDistance(v1)
	dist1to0 := v1.ManhattanDistance(v0)

	if dist0to1 != dist1to0 {
		t.Errorf("Distances are not the same: %f, %f", dist0to1, dist1to0)
	}
}

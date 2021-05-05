// A __very__ basic Spatial Hash Grid implementation
// Currently, a bucket size of 1x1 is used

package spatialhashgrid

import (
	"math"

	"github.com/tmhmitchell/ebitoolbox/datastructures/vector"
)

type Client interface {
	// X returns the origin of the Client on the X axis
	X() float64

	// Y returns the origin of the Client on the Y axis
	Y() float64

	Width() float64
	Height() float64
}

type SpatialHashGrid struct {
	// buckets is the underlying series of cells that our grid contains
	// A map is used to allow for the grid to be infinite, without needing to
	// write adding any logic requiring array re-sizing as would be needed with
	// a 2D array.
	buckets map[vector.Vec2][]Client
}

func New() *SpatialHashGrid {
	return &SpatialHashGrid{
		buckets: make(map[vector.Vec2][]Client),
	}
}

func (shg SpatialHashGrid) Length() int { return len(shg.buckets) }

// bounds returns the minimum (inclusive) and maximum (exclusive) bounds of the
// buckets a Client would be entered into, in the order x0, y0, x1, y1.
func bounds(c Client) (float64, float64, float64, float64) {
	return math.Floor(c.X()), math.Floor(c.Y()),
		math.Ceil(c.X() + c.Width()), math.Ceil(c.Y() + c.Height())

}

func (shg *SpatialHashGrid) Insert(c Client) {
	x0, y0, x1, y1 := bounds(c)

	for y := y0; y < y1; y++ {
		for x := x0; x < x1; x++ {
			bk := vector.NewVec2(x, y)

			// XXX: This allocates a new array if there isn't one available,
			// which is nice from a usability standpoint but might impact
			// performance - perhaps we should avoid deleting them in .Remove?
			shg.buckets[bk] = append(shg.buckets[bk], c)
		}
	}
}

func (shg *SpatialHashGrid) Remove(c Client) {
	x0, y0, x1, y1 := bounds(c)

	for y := y0; y < y1; y++ {
		for x := x0; x < x1; x++ {
			bk := vector.NewVec2(x, y)

			// No point iterating over a non-existant bucket
			bucket, ok := shg.buckets[bk]
			if !ok {
				continue
			}

			// If the client we're removing isn't in the bucket, skip it
			present := false
			for _, bc := range bucket {
				if bc == c {
					present = true
					break
				}
			}
			if !present {
				continue
			}

			// If the bucket only has a single element, delete it to prevent empty
			// buckets being left behind. This might not be the best approach?
			if len(bucket) == 1 {
				delete(shg.buckets, bk)
				continue
			}

			nb := make([]Client, len(bucket)-1)
			i := 0

			// Re-construct the bucket, but without c
			for _, elem := range bucket {
				if elem == c {
					continue
				}

				nb[i] = elem
				i++
			}

			shg.buckets[bk] = nb
		}
	}
}

// ClientsIn returns an array of all the clients inside a given rectangle
func (shg SpatialHashGrid) ClientsIn(x0, y0, x1, y1 float64) []Client {
	// We use a map as a set to ensure a client isn't returned multiple times
	cs := make(map[Client]struct{})

	for y := y0; y < y1; y++ {
		for x := x0; x < x1; x++ {
			bk := vector.NewVec2(x, y)

			// No point iterating over a non-existant bucket
			bucket, ok := shg.buckets[bk]
			if !ok {
				continue
			}

			for _, c := range bucket {
				cs[c] = struct{}{}
			}
		}
	}

	// Convert our "set" into an array
	csa := make([]Client, len(cs))
	i := 0
	for c := range cs {
		csa[i] = c
		i++
	}

	return csa
}

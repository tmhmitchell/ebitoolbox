package spatialhashgrid_test

import (
	"testing"

	"github.com/tmhmitchell/ebitoolbox/datastructures/spatialhashgrid"
)

type TestEntity struct{ x, y, w, h float64 }

func (te TestEntity) X() float64      { return te.x }
func (te TestEntity) Y() float64      { return te.y }
func (te TestEntity) Width() float64  { return te.w }
func (te TestEntity) Height() float64 { return te.h }

func TestInsert(t *testing.T) {
	tt := []struct {
		client          spatialhashgrid.Client
		expectedBuckets int
	}{
		{TestEntity{0, 0, 1, 1}, 1},
		{TestEntity{0.5, 0.5, 1, 1}, 4},
	}

	for _, item := range tt {
		grid := spatialhashgrid.New()

		grid.Insert(item.client)

		if grid.Length() != item.expectedBuckets {
			t.Logf("Expected %d buckets, found %d\n", item.expectedBuckets, grid.Length())
			t.FailNow()
		}
	}
}

func TestDisjointRemove(t *testing.T) {
	// Edge case where removing a client would remove 1-elem buckets containing !client
	c0 := TestEntity{0, 0, 1, 1}
	c1 := TestEntity{0.5, 0.5, 1, 1}

	grid := spatialhashgrid.New()

	grid.Insert(c0)
	grid.Remove(c1)

	if grid.Length() != 1 {
		t.Logf("Expected 1 buckets, found %d\n", grid.Length())
		t.FailNow()
	}
}

func TestInsertAndRemove(t *testing.T) {
	grid := spatialhashgrid.New()
	c0 := TestEntity{1, 1, 1, 1}
	c1 := TestEntity{0, 0, 2, 2}

	grid.Insert(c0)
	if grid.Length() != 1 {
		t.Logf("Expected 1 buckets, found %d\n", grid.Length())
		t.FailNow()
	}

	grid.Insert(c1)
	if grid.Length() != 4 {
		t.Logf("Expected 4 buckets, found %d\n", grid.Length())
		t.FailNow()
	}

	grid.Remove(c0)
	if grid.Length() != 4 {
		t.Logf("Expected 4 buckets, found %d\n", grid.Length())
		t.FailNow()
	}

	grid.Remove(c1)
	if grid.Length() != 0 {
		t.Logf("Expected 0 buckets, found %d\n", grid.Length())
		t.FailNow()
	}

	grid.Remove(TestEntity{1, 1, 1, 1})
}

func TestClientsIn(t *testing.T) {

	tt := []struct {
		clients         []spatialhashgrid.Client
		x0, y0, x1, y1  float64
		expectedClients int
	}{
		{
			[]spatialhashgrid.Client{TestEntity{1, 1, 1, 1}},
			0, 0, 2, 2,
			1,
		},
		{
			[]spatialhashgrid.Client{TestEntity{0.5, 0.5, 1, 1}},
			0, 0, 2, 2,
			1,
		},
		// Ensure that non-integer boundaries are correctly adjusted
		{
			[]spatialhashgrid.Client{TestEntity{2, 2, 1, 1}},
			1.5, 1.5, 3.5, 3.5,
			1,
		},
	}

	for _, ti := range tt {
		grid := spatialhashgrid.New()

		for _, c := range ti.clients {
			grid.Insert(c)
		}

		retrieved := grid.ClientsIn(ti.x0, ti.y0, ti.x1, ti.y1)

		if len(retrieved) != ti.expectedClients {
			t.Logf(
				"Expected %d client, found %d\n",
				ti.expectedClients,
				len(retrieved),
			)
			t.FailNow()
		}
	}
}

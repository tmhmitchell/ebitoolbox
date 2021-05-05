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
		client          spatialhashgrid.Client
		expectedClients int
	}{
		{TestEntity{1, 1, 1, 1}, 1},
		{TestEntity{0.5, 0.5, 1, 1}, 1},
	}

	for _, item := range tt {
		grid := spatialhashgrid.New()
		grid.Insert(item.client)

		retrieved := grid.ClientsIn(0, 0, 2, 2)

		if len(retrieved) != item.expectedClients {
			t.Logf("Expected %d client, found %d\n", item.expectedClients, len(retrieved))
			t.FailNow()
		}
	}
}

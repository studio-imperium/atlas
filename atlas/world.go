package atlas

import (
	"sync"
	"math/rand"
)

type World struct {
	points []Point
	triangles []Triangle
	cellByOrigin map[Point]*Cell
	
	Cells []Cell `json:"cells"`
	Size int `json:"size"`
}

func assignCol(x int, world *World) {
	for y := 0; y < world.Size; y++ {
		tile := Tile{
			X: x,
			Y: y,
			Value: 0,
		}
		cell := world.GetNearestCell(tile.point())
		cell.addTile(tile)
	}
}

func newWorld(density int, size int, seed int64) *World {
	rnd := rand.New(rand.NewSource(seed))
	
	world := new(World)
	world.points = make([]Point, size)
	world.triangles = createTriangles(world.points)
	world.cellByOrigin = make(map[Point]*Cell)
	for i := range world.points {
		world.points[i] = Point{
			X: float64(size) * rnd.Float64(),
			Y: float64(size) * rnd.Float64(),
		}
		
		cell := NewCell(world.points[i])
		world.cellByOrigin[cell.Origin] = &cell
		world.Cells = append(world.Cells, cell)
	}
	
	// Assign vertices to each cell
	//for _, triangle := range world.triangles {
		//vertex := NewVertex(triangle, cellByOrigin)
		
		//for _, origin := range triangle.Points {
			//cell := world.cellByOrigin[origin]
			//cell.vertices = append(cell.vertices, vertex)
		//}
	//}
	
	var wg sync.WaitGroup
	wg.Add(world.Size)
	for x := 0; x < world.Size; x++ {
		for y := 0; y < world.Size; y++ {
			tile := Tile{
				X: x,
				Y: y,
				Value: 0,
			}
			cell := world.GetNearestCell(tile.point())
			cell.addTile(tile)
		}
	}
	wg.Wait()
	
	return world
}

func (world *World) GetNearestCell(pt Point) *Cell {
	nearest := world.Cells[0].Origin
	nearestDist := distance(nearest, pt)
	
	for idx := range world.Cells {
		cell := &world.Cells[idx]
		cellDist := distance(cell.Origin, pt)
		if nearestDist > cellDist {
			nearest = cell.Origin
			nearestDist = cellDist
		}
	}
	
	return world.cellByOrigin[nearest]
}

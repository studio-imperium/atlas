package atlas

type Tile struct {
	x int32
	y int32
	value int8
}

type World struct {
	cellTiles map[int32][]Tile
	tileCells [][]int32
	
	points []Point
	triangles []Triangle
}

func (world *World) Populate() {
	
}

func GenerateWorld(size float64) *World {
	return new(World)
}

func GenerateData(n int, radius float64, seed int64) []Triangle {
	points := generatePoints(n, radius, seed)
	triangles := CreateTriangles(points)
	return triangles
}





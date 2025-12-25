package atlas

const DENSITY_LOW = 10
const DENSITY_MED = 50
const DENSITY_HIGH = 200

type Tile struct {
	X int `json:"x"`
	Y int `json:"y"`
	Value int8 `json:"value"`
}

func (tile Tile) getX() float64 {
	return float64(tile.X)
}

func (tile Tile) getY() float64 {
	return float64(tile.Y)
}

func NewTile(x int, y int) Tile {
	return Tile{
		X: x,
		Y: y,
	}
}


type Cell struct {
	Origin Point
	Tiles []Tile `json:"tiles"`
}

func (cell *Cell) addTile(tile Tile) {
	cell.Tiles = append(cell.Tiles, tile)
}

func NewCell(origin Point) Cell {
	return Cell{
		Origin: origin,
		Tiles: []Tile{},
	}
}

type World struct {
	points []Point
	triangles []Triangle
	
	Cells []Cell `json:"cells"`
	Size int `json:"size"`
}

func (world *World) PrintPoints() []Point {
	return world.points
}

func (world *World) GetNearestCell(tile Vector) *Cell {
	nearest := &world.Cells[0]
	for idx := range world.Cells {
		cell := &world.Cells[idx]
		if distance(nearest.Origin, tile) > distance(cell.Origin, tile) {
			nearest = cell
		}
	}
	
	return nearest
}

func (world *World) Populate(density int) {
	world.points = generatePoints(density, world.Size, 11)
	world.triangles = CreateTriangles(world.points)
	
	for _, point := range world.points {
		world.Cells = append(world.Cells, NewCell(point))
	}
	
	for x := 0; x < world.Size; x++ {
		for y := 0; y < world.Size; y++ {
			tile := NewTile(x, y)
			cell := world.GetNearestCell(tile)
			cell.addTile(tile)
		}
	}
}

func GenerateData(n int, radius int, seed int64) []Triangle {
	points := generatePoints(n, radius, seed)
	triangles := CreateTriangles(points)
	return triangles
}

func GenerateWorld(size int) *World {
	world := new(World)
	world.Size = size
	world.Populate(DENSITY_HIGH)
	return world
}




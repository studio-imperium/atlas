package atlas

type Cell struct {
	vertices []Vertex
	
	Origin Point
	Tiles []Tile `json:"tiles"`
}

func NewCell(origin Point) Cell {
	return Cell{
		vertices: []Vertex{},
		Origin: origin,
		Tiles: []Tile{},
	}
}

func (cell *Cell) addTile(tile Tile) {
	cell.Tiles = append(cell.Tiles, tile)
}

func (cell *Cell) GetAdjacentCells() []*Cell {
	seen := make(map[Point]bool)
	seen[cell.Origin] = true
	cells := []*Cell{}
	for _, vertex := range cell.vertices {
		for _, adjacentCell := range vertex.cells {
			if !seen[adjacentCell.Origin] {
				cells = append(cells, adjacentCell)
				seen[adjacentCell.Origin] = true
			}
		}
	}
	return cells
}

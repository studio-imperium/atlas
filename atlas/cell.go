package atlas

import "sync"

type Cell struct {
	vertices []Vertex
	value int8
	
	Origin Point
	Tiles []Tile `json:"tiles"`
	
	mu sync.Mutex
}

func NewCell(origin Point) *Cell {
	return &Cell{
		value: 0,
		vertices: []Vertex{},
		Origin: origin,
		Tiles: []Tile{},
	}
}

func (cell *Cell) setValue(value int8) {
	cell.value = value
	for idx := range cell.Tiles {
		tile := &(cell.Tiles[idx])
		tile.Value = value
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

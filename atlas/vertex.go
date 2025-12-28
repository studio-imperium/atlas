package atlas

type Vertex struct {
	location Point
	cells [3]*Cell
}

func NewVertex(triangle Triangle, cells map[Point]*Cell) Vertex {
	return Vertex{
		location: triangle.center,
		cells: [3]*Cell{
			cells[triangle.Points[0]],
			cells[triangle.Points[1]],
			cells[triangle.Points[2]],
		},
	}
}

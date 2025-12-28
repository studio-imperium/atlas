package atlas

import "math"

type Triangle struct {
	Points [3]Point `json:"points"`
	center Point `json:"center"`
	radius float64 `json:"radius"`
}

func newTriangle(pts [3]Point) Triangle {
	triangle := new(Triangle)
	triangle.Points = pts
	triangle.center = circumcenter(pts)
	triangle.radius = distance(triangle.center, pts[0])
	
	return *triangle
}


func (triangle Triangle) includesPoint(pt Point) bool {
	for _, _pt := range triangle.Points {
		if _pt.X == pt.X && _pt.Y == pt.Y {
			return true
		}
	}
	return false
}

func (triangle Triangle) withinCircumcircle(pt Point) bool {
	return distance(pt, triangle.center) <= triangle.radius
}

func (triangle Triangle) validDelauney(pts *[]Point) bool {
	for _, pt := range *pts {
		if triangle.withinCircumcircle(pt) && !triangle.includesPoint(pt) {
			return false
		}
	}
	return true
}

func (triangle Triangle) reform(pt Point, pts *[]Point) []Triangle {
	validTriangles := []Triangle{}
	triangles := []Triangle{
		newTriangle([3]Point{
			pt,
			triangle.Points[0],
			triangle.Points[1],
		}),
		newTriangle([3]Point{
			pt,
			triangle.Points[1],
			triangle.Points[2],
		}),
		newTriangle([3]Point{
			pt,
			triangle.Points[2],
			triangle.Points[0],
		}),
	}
	
	for _, triangle := range triangles {
		if triangle.validDelauney(pts) {
			validTriangles = append(validTriangles, triangle)
		}
	}
	
	return validTriangles
}

func circumcenter(pts [3]Point) Point {
	p1 := pts[0]
	p2 := pts[1]
	p3 := pts[2]
	
	a := -(p2.X - p1.X) / (p2.Y - p1.Y)
	c := -(p3.X - p2.X) / (p3.Y - p2.Y)
    b := ((p2.Y + p1.Y) / 2) - (a * (p2.X + p1.X) / 2)
    d := ((p3.Y + p2.Y) / 2) - (c * (p3.X + p2.X) / 2)

    x := (d - b) / (a - c)
    y := a * x + b

    if math.IsNaN(x) || math.IsNaN(y) {
        return circumcenter([3]Point{p3, p1, p2})
    }
	return Point{
		X: x,
		Y: y,
	}
}

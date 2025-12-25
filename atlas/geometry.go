package atlas

import (
	"math"
)

type Vector interface {
	getX() float64
	getY() float64
}

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func (point Point) getX() float64 {
	return point.X
}
func (point Point) getY() float64 {
	return point.Y
}

func distance(vector1 Vector, vector2 Vector) float64 {
	return math.Pow(math.Pow(vector1.getX() - vector2.getX(), 2) + math.Pow(vector1.getY() - vector2.getY(), 2), 0.5)
}

func (point1 Point) add(point2 Point) Point {
	return Point{
		X: point1.X + point2.X,
		Y: point1.Y + point2.Y,
	}
}

func (point1 Point) subtract(point2 Point) Point {
	return Point{
		X: point1.X - point2.X,
		Y: point1.Y - point2.Y,
	}
}

func (point1 Point) multiply(point2 Point) Point {
	return Point{
		X: point1.X * point2.X,
		Y: point1.Y * point2.Y,
	}
}

func (point Point) sqr() Point {
	return Point{
		X: math.Pow(point.X, 2),
		Y: math.Pow(point.Y, 2),
	}
}

func (point1 Point) divide(point2 Point) Point {
	return Point{
		X: point1.X / point2.X,
		Y: point1.Y / point2.Y,
	}
}

func NewPoint(x float64, y float64) Point {
	return Point{
		X: x,
		Y: y,
	}
}



type Triangle struct {
	Points [3]Point `json:"points"`
	center Point `json:"center"`
	radius float64 `json:"radius"`
}

func (triangle Triangle) IncludesPoint(point Point) bool {
	for _, p := range triangle.Points {
		if p.X == point.X && p.Y == point.Y {
			return true
		}
	}
	return false
}

func (triangle Triangle) withinCircumcircle(point Point) bool {
	return distance(point, triangle.center) <= triangle.radius
}

func (triangle Triangle) validDelauney(points *[]Point) bool {
	for _, point := range *points {
		if triangle.withinCircumcircle(point) && !triangle.IncludesPoint(point) {
			return false
		}
	}
	return true
}

func (triangle Triangle) reform(point Point, points *[]Point) []Triangle {
	validTriangles := []Triangle{}
	triangles := []Triangle{
		NewTriangle([3]Point{
			point,
			triangle.Points[0],
			triangle.Points[1],
		}),
		NewTriangle([3]Point{
			point,
			triangle.Points[1],
			triangle.Points[2],
		}),
		NewTriangle([3]Point{
			point,
			triangle.Points[2],
			triangle.Points[0],
		}),
	}
	
	for _, triangle := range triangles {
		if triangle.validDelauney(points) {
			validTriangles = append(validTriangles, triangle)
		}
	}
	
	return validTriangles
}

func circumcenter(points [3]Point) Point {
	p1 := points[0]
	p2 := points[1]
	p3 := points[2]
	
	a := -(p2.X - p1.X) / (p2.Y - p1.Y)
	c := -(p3.X - p2.X) / (p3.Y - p2.Y)
    b := ((p2.Y + p1.Y) / 2) - (a * (p2.X + p1.X) / 2)
    d := ((p3.Y + p2.Y) / 2) - (c * (p3.X + p2.X) / 2)

    x := (d - b) / (a - c)
    y := a * x + b

    if math.IsNaN(x) || math.IsNaN(y) {
        return circumcenter([3]Point{p2, p3, p1})
    }
	return Point{
		X: x,
		Y: y,
	}
}


func NewTriangle(points [3]Point) Triangle {
	triangle := new(Triangle)
	triangle.Points = points
	triangle.center = circumcenter(points)
	triangle.radius = distance(triangle.center, points[0])
	
	return *triangle
}

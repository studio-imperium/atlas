package atlas

import (
	"math"
	"fmt"
)

type Point struct {
	x float64
	y float64
}

func distance(point1 Point, point2 Point) float64 {
	return math.Pow(math.Pow((point1.x - point2.x), 2) +
			math.Pow((point1.y - point2.y), 2), 0.5)
}

func (point1 Point) add(point2 Point) Point {
	return Point{
		x: point1.x + point2.x,
		y: point1.y + point2.y,
	}
}

func (point1 Point) subtract(point2 Point) Point {
	return Point{
		x: point1.x - point2.x,
		y: point1.y - point2.y,
	}
}

func (point1 Point) multiply(point2 Point) Point {
	return Point{
		x: point1.x * point2.x,
		y: point1.y * point2.y,
	}
}

func (point Point) sqr() Point {
	return Point{
		x: math.Pow(point.x, 2),
		y: math.Pow(point.y, 2),
	}
}

func (point1 Point) divide(point2 Point) Point {
	return Point{
		x: point1.x / point2.x,
		y: point1.y / point2.y,
	}
}

func NewPoint(x float64, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}



type Triangle struct {
	points [3]Point
	center Point
	radius float64
}

func (triangle Triangle) IncludesPoint(point Point) bool {
	for _, p := range triangle.points {
		if p.x == point.x && p.y == point.y {
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
			triangle.points[0],
			triangle.points[1],
		}),
		NewTriangle([3]Point{
			point,
			triangle.points[1],
			triangle.points[2],
		}),
		NewTriangle([3]Point{
			point,
			triangle.points[2],
			triangle.points[0],
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
	fmt.Println(points)
	
	p1 := points[0]
	p2 := points[1]
	p3 := points[2]
	
	a := -(p2.x - p1.x) / (p2.y - p1.y)
	c := -(p3.x - p2.x) / (p3.y - p2.y)
    b := ((p2.y + p1.y) / 2) - (a * (p2.x + p1.x) / 2)
    d := ((p3.y + p2.y) / 2) - (c * (p3.x + p2.x) / 2)

    x := (d - b) / (a - c)
    y := a * x + b

    if math.IsNaN(x) || math.IsNaN(y) {
        return circumcenter([3]Point{p2, p3, p1})
    }
	return Point{
		x: x,
		y: y,
	}
}


func NewTriangle(points [3]Point) Triangle {
	triangle := new(Triangle)
	triangle.points = points
	triangle.center = circumcenter(points)
	triangle.radius = distance(triangle.center, points[0])
	
	return *triangle
}

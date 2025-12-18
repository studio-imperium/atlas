package atlas

import (
	"math"
	"math/rand"
)

func generatePoints(n int, radius float64, seed int64) []Point {
	r := rand.New(rand.NewSource(seed))
	points := []Point{}
	
	for i := 0; i < n; i++ {
		points = append(points, Point{
			x: radius * r.Float64(),
			y: radius * r.Float64(),
		})
	}
	
	return points
}

func addPoint(triangles []Triangle, point Point, points *[]Point) []Triangle {
	newTriangles := []Triangle{}
	
	for _, triangle := range triangles {
		if triangle.withinCircumcircle(point) {
			newTriangles = append(newTriangles, triangle.reform(point, points)...)
		} else {
			newTriangles = append(newTriangles, triangle)
		}
	}
	
	return newTriangles
}

func CreateTriangles(points []Point) []Triangle {
	maxY := points[0].y
	maxX := points[0].x
	for _, point := range points {
		maxY = math.Max(maxY, point.y)
		maxX = math.Max(maxX, point.x)
	}
	
	addedPoints := []Point{
		NewPoint(0,0),
		NewPoint(maxX + 1,0),
		NewPoint(0,maxY + 1),
		NewPoint(maxX + 1,maxY + 1),
	}
	
	triangles := []Triangle{
		NewTriangle([3]Point{
			addedPoints[0],
			addedPoints[1],
			addedPoints[2],
		}),
		NewTriangle([3]Point{
			addedPoints[1],
			addedPoints[2],
			addedPoints[3],
		}),
	}
	
	for _, point := range points {
		triangles = addPoint(triangles, point, &addedPoints)
		addedPoints = append(addedPoints, point)
	}

	return triangles
}

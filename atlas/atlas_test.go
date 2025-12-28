package atlas

import (
	"testing"
	"fmt"
)

func TestTriangles(t * testing.T) {
	triangles := GenerateWorld(10)
	
	t.Errorf("Triangles are: %s", fmt.Sprintln(triangles))
}

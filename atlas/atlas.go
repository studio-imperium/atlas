package atlas

const DENSITY_LOW = 10
const DENSITY_MED = 100
const DENSITY_HIGH = 500

func GenerateWorld() *World {
	size := 1000
	world := newWorld(DENSITY_HIGH, size, 100)
	return world
}




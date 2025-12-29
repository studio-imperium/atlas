package atlas

func NewWorld(size int, density int, seed int64) *World {
	world := newWorld(size, density, seed)
	return world
}

func TemplateWorld(size int) *World {
	world := newWorld(2 * size, size, 5)
	world.infect(world.Cells[0], 0.5)
	world.infect(world.Cells[5], 0.5)
	
	return world
}



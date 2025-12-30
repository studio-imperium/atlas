package atlas

import "fmt"

func NewWorld(size int, density int, seed int64) *World {
	world := newWorld(size, density, seed)
	return world
}

//"deepWater" : 0,
//"water" : 1,
//"sand" : 2,
//"grass" : 3,
//"sandstone" : 4,
//"dryGrass" : 5,
//"coals" : 6,
//"ruins" : 7,
//"snow" : 8,

func TemplateWorld(size int) *World {
	biomes := []Biome{
		Biome{
			[]Modifier{
				NewVoronoi(20, 2,4,4,5),
				NewSelectiveBorder(4,5),
			},
		},
		Biome{
			[]Modifier{
				NewVoronoi(10, 6,6,7,7,7,7,8),
				NewSelectiveBorder(7,6),
			},
		},
		Biome{
			[]Modifier{
				NewVoronoi(30, 6,6,6,6,7,7,7,7,7,7,7,8),
				NewSelectiveBorder(7,6),
				NewBorder(8),
			},
		},
	}
	
	world := newWorld(2 * size, size, 5)
	world.infect(biomes, 0.3)
	world.infect(biomes, 0.2)
	
	fmt.Println(world.Cells[0].biome)
	
	return world
}

func NewBiomes(biomes int8) []Biome {
	biome := make([]Biome, biomes)
	
	for idx := range biome {
		biome[idx] = Biome{[]Modifier{NewBase()}}
	}
	
	return biome
}


package atlas

func NewWorld(size int, density int, seed int64) *World {
	world := newWorld(size, density, seed)
	return world
}

func NewTemplateWorld(size int) *World {
	world := newWorld(size, size, 0)
	return world
}


// Tiles
var DEEPWATER int8 = 0
var WATER int8 = 1
var GRASS int8 = 2
var STONE int8 = 3
var SAND int8 = 4
var SANDSTONE int8 = 5
var DRYGRASS int8 = 6
var RUBBLE int8 = 7
var DARKSTONE int8 = 8
var SNOW int8 = 9
var ICE int8 = 10

var OceanMap []Biome = []Biome{
	NewBiome(
		NewFill(DEEPWATER),
	),
	NewBiome(
		NewFill(WATER),
	),
}

var IslandMap []Biome = []Biome{
	NewBiome(
		NewFill(DEEPWATER),
	),
	NewBiome(
		NewFill(WATER),
	),
	NewBiome(
		NewFill(SAND),
	),
	NewBiome(
		NewFill(GRASS),
	),
	NewBiome(
		NewFill(GRASS),
	),
	NewBiome(
		NewFill(GRASS),
	),
}

var DesertMountainsMap []Biome = []Biome{
	NewBiome(
		NewPattern(27, RUBBLE, SNOW),
		NewSelectiveBorder(SNOW, ICE),
		NewSelectiveBorder(DARKSTONE, RUBBLE),
		NewBorder(ICE),
		NewSelectiveExternalBorder(SNOW, ICE),
	),
	NewBiome(
		NewVoronoi(40, RUBBLE,DARKSTONE,DARKSTONE,SNOW,ICE),
		NewSelectiveBorder(DARKSTONE,SNOW),
		NewSelectiveBorder(SNOW,ICE),
		NewBorder(SNOW),
	),
	NewBiome(
		NewVoronoi(40, RUBBLE,RUBBLE,DARKSTONE,DARKSTONE,DARKSTONE,DARKSTONE,DARKSTONE,SNOW),
		NewSelectiveBorder(DARKSTONE,SNOW),
		NewBorder(SNOW),
	),
	NewBiome(
		NewVoronoi(40, RUBBLE,RUBBLE,DARKSTONE,DARKSTONE,SNOW),
		NewSelectiveBorder(DARKSTONE,SNOW),
		NewBorder(DARKSTONE),
	),
	NewBiome(
		NewCropCircle(40, DRYGRASS,SAND),
		NewSelectiveBorder(SANDSTONE,SAND),
		NewSelectiveBorder(SAND,DRYGRASS),
		NewSelectiveBorder(SAND,DRYGRASS),
		NewSelectiveBorder(SAND,DRYGRASS),
		NewSelectiveBorder(SANDSTONE,SAND),
	),
	NewBiome(
		NewVoronoi(40, SAND,SANDSTONE,DRYGRASS),
		NewSelectiveBorder(SANDSTONE,DRYGRASS),
	),
	NewBiome(
		NewPattern(5, DRYGRASS,SAND),
		NewSelectiveBorder(SANDSTONE,SAND),
	),
	NewBiome(
		NewPattern(3.2, GRASS,STONE),
		NewBorder(STONE),
	),
	NewBiome(
		NewCropCircle(20, WATER,GRASS),
		NewSelectiveBorder(STONE,WATER),
	),
}

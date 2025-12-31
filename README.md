# Atlas map generator

[![GoDoc](https://godoc.org/github.com/studio-imperium/atlas?status.svg)](https://godoc.org/github.com/studio-imperium/atlas)

Atlas is a 2d map generator with a unique twist.
Most tilemaps are 2d arrays, with tiles[y][x] being the tile at (x,y).
World instead has a Cells property, each cell containing their respective tiles in Cell.Tiles.

# Starting

When creating a world with NewWorld, the density parameter specifies how many cells you want to split the world into.

To create a 100x100 world with 100 cells and all tiles set to a value of 0, NewWorld(100, 100, seed int64) (where seed can be any number)
To create a 100x100 world with 100 cells and all tiles set to a value of 0, alternatively use NewTemplateWorld(100) (size = density)

Cells are intended to be used as chunks and be sent to players as seen in the example (examples/mapviewer/main.go). Though it might seem excessive compared to regular chunks, finding a players cell is O(n) with respect to the density of the world, and finding adjacent cells to the player is O(1) since they are cached while generating the world.

Voronoi cell based worlds serve another function, generating terrain becomes beautiful since you can apply functions to modify all the tiles in a cell locally. The intuition is we have O(1) adjacent cell access, so a fast way to generate terrain would be to use an "infection" method, choosing our patient-0 cell and spreading biomes with some parameters to change how the spread behaves.

# Terrain

Since I built this for Kingdom Crushers, I set up a World.Infect method which takes a list of "biomes" (functions which modify a cells tiles) and the "decay" (strength of infection) to create a modular terrain generation system.

Modifier: Method which modifies a cell's tiles
Biome: Struct with a []Modifier property, applies each Modifier in order.
Infection method: World.Infect(biomes []Biome, decay float64)

Infection assigns "patient-0" the first biome, and as it spreads to the adjacent cells it goes through the Biomes first through last.
So if you had biomes = []{ Biome1, Biome2, Biome3 },
Biome1 is first, then Biome2 would spread for a bit, then eventually once it dies off Biome3 would fill the rest of the cells in the world.

There are some built in Modifiers, including:

- NewFill(value int8) generates a modifier which fills all of a cells tiles with some int8.
- NewBorder(value int8) replaces all the edge tiles of a cell with some int8.
- NewVoronoi(density int, values ...int8) creates a mini voronoi diagram inside the cell, with ${density} cells, each filled with a random value from values.

The most powerful ones are the pattern, cropcircles & selective border.
You can make new biomes with NewBiome(mods ...Modifier), where the order of your modifiers will define their execution order on the cell.

# Example

Say you want to generate a 100x100 world which has every tile as 0:
world := NewTemplateWorld(100)

Now, if you want to make all the tiles have a value of 1:
    flatMap := []Biome{
        NewBiome(NewFill(1))
    }
    world.Infect(flatMap, 0)

Say instead you want some with a value of 1 and some with a value of 0 in polygons:
    flatMap := []Biome{
        NewBiome(NewVoronoi(10, 0,1))
    }
    world.Infect(flatMap, 0)

Lets assign meaning to the numbers as we would in a real use case.
Now if we want to generate a Island,
    
    island := NewTemplateWorld(100)
    islandMap := []Biome{
        NewBiome(NewFill(GRASS)),
        NewBiome(NewFill(GRASS)),
        NewBiome(NewFill(GRASS)),
        NewBiome(NewFill(SAND)),
        NewBiome(NewFill(WATER)),
    }
    world.Infect(islandMap, 0.5)





### Installation

    go get github.com/studio-imperium/atlas

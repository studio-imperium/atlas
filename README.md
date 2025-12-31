# Atlas map generator

[![GoDoc](https://godoc.org/github.com/studio-imperium/atlas?status.svg)](https://godoc.org/github.com/studio-imperium/atlas)

Atlas is a 2d map generator with a unique twist.

Most tilemaps are 2d arrays, with tiles[y][x] being the tile at (x,y).

World instead has a Cells property, each cell containing their respective tiles in Cell.Tiles.

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

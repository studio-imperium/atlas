# Atlas map generator

[![GoDoc](https://godoc.org/github.com/studio-imperium/atlas?status.svg)](https://godoc.org/github.com/studio-imperium/atlas)

Atlas is a 2d map generator with a unique twist.

Most tilemaps are 2d arrays, with tiles[y][x] being the tile at (x,y).

We instead use cells, each cell containing their respective tiles in Cell.Tiles.

[Read more](https://williamqm.com/writing/mapgen/)

### Examples

[Live demo](http://100.28.2.231:6001/)
![map1](https://williamqm.com/writing/mapgen/map11.png)
[Live demo](http://100.28.2.231:2000/)
![map1](https://williamqm.com/writing/mapgen/map12.png)
[Live demo](http://100.28.2.231:1000/)
![map1](https://williamqm.com/writing/mapgen/map13.png)



### Installation

    go get github.com/studio-imperium/atlas

### Choosing the biome origin

`Infect` starts at a random cell using the world's seed. Use `InfectFrom` to
start at the cell nearest a specific point, such as the center of an island:

```go
world := atlas.NewWorld(256, 100, 11)
biomes := []atlas.Biome{
    atlas.NewBiome(atlas.NewFill(2)),
    atlas.NewBiome(atlas.NewFill(6)),
    atlas.NewBiome(atlas.NewFill(1)),
}
center := atlas.Point{X: float64(world.Size / 2), Y: float64(world.Size / 2)}
world.InfectFrom(biomes, 0.25, center)
```

`InfectFrom` uses the same spread and modifier rules as `Infect`, without
advancing the world's random source. `cell.GetBiome()` returns the assigned
biome index as `int8`; tile values also remain `int8`.

For custom map exports, use each cell's position in `world.Cells` as its ID
and place its tiles at `tile.X + tile.Y*world.Size` in a flat array.

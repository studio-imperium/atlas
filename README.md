# Atlas map generator

[![GoDoc](https://godoc.org/github.com/studio-imperium/atlas?status.svg)](https://godoc.org/github.com/studio-imperium/atlas)

Atlas is a 2d map generator with a unique twist.

Most tilemaps are 2d arrays, with tiles[y][x] being the tile at (x,y).

World instead has a Cells property, each cell containing their respective tiles in Cell.Tiles.

[Read more](https://williamqm.com/writing/mapgen/)

### Examples

![map1](https://williamqm.com/writing/mapgen/map11.png)
![map1](https://williamqm.com/writing/mapgen/map12.png)
![map1](https://williamqm.com/writing/mapgen/map13.png)



### Installation

    go get github.com/studio-imperium/atlas

package atlas

import (
	"math"
	"math/rand"
	"sync"
)

type Modifier func(*Cell)

type Biome struct {
	modifiers []Modifier
}

func (biome *Biome) SetModifier(idx int, mod Modifier) {
	biome.modifiers[idx] = mod
}

func (world *World) infect(biomes []Biome, decay float64) {
	
	// Get "patient 0"
	idx := world.rnd.Int() % len(world.Cells)
	origin := world.Cells[idx]
	
	// Keep track of who to infect
	changedCells := []*Cell{}
	seen := make(map[*Cell]bool)
	queue := []*Cell{
		origin,
	}
	
	// Keep track of biome
	biomeCount := len(biomes)
	currentBiome := float64(biomeCount)
	convertBiome := func() int8 {
		return int8(math.Floor(currentBiome))
	}
	
	// Infect
	for len(queue) > 0 {
		currentBiome -= decay
		currentBiome = math.Max(0, currentBiome)
		temp := []*Cell{}
		
		for _, cell := range queue {
			// Assign biome to current cell
			biomeInt8 := convertBiome()
			if cell.biome <= biomeInt8 {
				changedCells = append(changedCells, cell)
				cell.biome = biomeInt8
			}
			
			// Infect adjacent cells
			for _, adj := range cell.GetAdjacentCells() {
				if !seen[adj] {
					seen[adj] = true
					temp = append(temp, adj)
				}
			}
		}
		
		queue = temp
	}
	
	// Now that biomes are assigned use the biome modifiers on each cell
	wg := &sync.WaitGroup{}
	wg.Add(len(changedCells))
	for _, cell := range changedCells {
		biome := biomes[cell.biome]
		go func() {
			defer wg.Done()
			for _, mod := range biome.modifiers {
				cell.mu.Lock()
				mod(cell)
				cell.mu.Unlock()
			}
		}()
	}
	wg.Wait()
}



// Modifiers

func NewBase() Modifier {
	return func(cell *Cell) {
		for idx := range cell.Tiles {
			tile := &(cell.Tiles[idx])
			tile.Value = cell.biome
		}
	}
}

func NewVoronoi(density int, values ...int8) Modifier {	
	return func(cell *Cell) {
		rnd := rand.New(rand.NewSource(int64(density)))
		origins := []Point{}
		getPoint := func() Point {
			return cell.Tiles[rnd.Int() % len(cell.Tiles)].point()
		}
		getValue := func(seed int) int8 {
			return values[seed % len(values)]
		}
		findNearest := func(pt Point) int8 {
			nearest := origins[0]
			nearestDist := distance(nearest, pt)
			
			for idx := range origins {
				origin := origins[idx]
				originDist := distance(origin, pt)
				if nearestDist > originDist {
					nearest = origin
					nearestDist = originDist
				}
			}
			
			return getValue(int(nearest.X + nearest.Y))
		}
		
		for idx := 0; idx < density; idx++ {
			pt := getPoint()
			origins = append(origins, pt)
		}
		
		for idx := range cell.Tiles {
			tile := &(cell.Tiles[idx])
			tile.Value = findNearest(tile.point())
		}
	}
}

func NewBorder(border int8) Modifier {
	return func(cell *Cell) {
		isBorder := func(cell *Cell, tile *Tile) bool {
			x := tile.X
			y := tile.Y
			adj := [4]Point{
				newPoint(x - 1, y),
				newPoint(x + 1, y),
				newPoint(x, y + 1),
				newPoint(x, y - 1),
			}
			for _, pt := range adj {
				if cell.grid[pt] == nil {
					return true
				}
			}
			
			return false
		}
		
		for idx := range cell.Tiles {
			tile := &(cell.Tiles[idx])
			
			if isBorder(cell, tile) {
				tile.Value = border
			}
		}
	}
}

func NewSelectiveBorder(border int8, around int8,) Modifier {
	return func(cell *Cell) {
		isBorder := func(cell *Cell, tile *Tile) bool {
			x := tile.X
			y := tile.Y
			adj := [4]Point{
				newPoint(x - 1, y),
				newPoint(x + 1, y),
				newPoint(x, y + 1),
				newPoint(x, y - 1),
			}
			for _, pt := range adj {
				if cell.grid[pt] == nil || cell.grid[pt].Value != around {
					return true
				}
			}
			
			return false
		}
		
		toChange := []*Tile{}
		for idx := range cell.Tiles {
			tile := &(cell.Tiles[idx])
			
			if tile.Value == around && isBorder(cell, tile) {
				toChange = append(toChange, tile)
			}
		}
		
		for _, tile := range toChange {
			tile.Value = border
		}
	}
}

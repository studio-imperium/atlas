package atlas

func (world *World) infect(patient0 *Cell, decay float64) {
	value := 8.0
	infected := make(map[*Cell]bool)
	queue := []*Cell{
		patient0,
	}
	
	for len(queue) > 0 {
		value -= decay
		tempQueue := []*Cell{}
		
		for _, cell := range queue {
			if cell.value < int8(value) {
				cell.setValue(int8(value))
			}
			for _, adj := range cell.GetAdjacentCells() {
				if world.rnd.Float64() < value && !infected[adj] {
					infected[adj] = true
					tempQueue = append(tempQueue, adj)
				}
			}
		}
		
		queue = tempQueue
	}
}

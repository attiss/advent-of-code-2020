package treemap

type TreeMap [][]bool

func (tm TreeMap) CountPathTrees(right, down int) int {
	trees := 0
	for x, y := 0, 0; y <= len(tm)-1; x, y = x+right, y+down {
		if x > len(tm[y])-1 {
			x %= len(tm[y])
		}

		if tm[y][x] {
			trees++
		}
	}
	return trees
}

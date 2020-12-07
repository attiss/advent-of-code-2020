package plane_seat

type PlaneSeat string

func (ps PlaneSeat) GetSeatID(rows, columns int) int {
	currentMinRow := 0
	currentMaxRow := rows - 1
	row := 0

	currentMinColumn := 0
	currentMaxColumn := columns - 1
	column := 0

	devideRange := func(low, high int, chooseLowerHalf bool) (newLow, newHigh int) {
		rowRange := (high - low + 1) / 2
		if chooseLowerHalf {
			if rowRange == 1 {
				newLow = low
				newHigh = low
			} else {
				newLow = low
				newHigh = high - rowRange
			}
		} else {
			if rowRange == 1 {
				newLow = high
				newHigh = high
			} else {
				newLow = low + rowRange
				newHigh = high
			}
		}
		return
	}

	for _, c := range ps {
		switch string(c) {
		case "F":
			currentMinRow, currentMaxRow = devideRange(currentMinRow, currentMaxRow, true)
		case "B":
			currentMinRow, currentMaxRow = devideRange(currentMinRow, currentMaxRow, false)
		case "L":
			currentMinColumn, currentMaxColumn = devideRange(currentMinColumn, currentMaxColumn, true)
		case "R":
			currentMinColumn, currentMaxColumn = devideRange(currentMinColumn, currentMaxColumn, false)
		}

		if currentMinRow == currentMaxRow {
			row = currentMinRow
		}
		if currentMinColumn == currentMaxColumn {
			column = currentMinColumn
		}
	}

	return row*8 + column
}

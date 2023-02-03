package utils

type Table []rune

func GenerateTable(key string) Table {
	var table Table
	used := make(map[rune]bool)

	for _, char := range key {
		table.Fill(char, used)
	}

	for char := 'A'; char <= 'Z'; char++ {
		table.Fill(char, used)
	}

	return table
}

func (table *Table) Fill(char rune, used map[rune]bool) {
	if used[char] {
		return
	}
	if char == 'J' {
		return
	}
	used[char] = true
	*table = append(*table, char)
}

func (table *Table) GetIndex(char1, char2 rune) (int, int) {
	i1, i2 := -1, -1
	for index, c := range *table {
		if c == char1 {
			i1 = index
		} else if c == char2 {
			i2 = index
		}
	}
	return i1, i2
}

func (table *Table) IsSameRow(i1, i2 int) bool {
	return i1/5 == i2/5
}

func (table *Table) IsSameColumn(i1, i2 int) bool {
	return i1%5 == i2%5
}

func (table *Table) ShiftHorizontal(i1, i2 int, n int) (int, int) {
	row1, col1 := table.convertIndexToRowAndColumn(i1)
	row2, col2 := table.convertIndexToRowAndColumn(i2)

	i1 = table.convertRowAndColumnToIndex(row1, (col1+n)%5)
	i2 = table.convertRowAndColumnToIndex(row2, (col2+n)%5)

	if (col1+n)%5 >= 0 {
		i1 = table.convertRowAndColumnToIndex(row1, (col1+n)%5)
	} else {
		i1 = table.convertRowAndColumnToIndex(row1, 5+(col1+n)%5)
	}

	if (col2+n)%5 >= 0 {
		i2 = table.convertRowAndColumnToIndex(row2, (col2+n)%5)
	} else {
		i2 = table.convertRowAndColumnToIndex(row2, 5+(col2+n)%5)
	}

	return i1, i2
}

func (table *Table) ShiftVertical(i1, i2 int, n int) (int, int) {
	row1, col1 := table.convertIndexToRowAndColumn(i1)
	row2, col2 := table.convertIndexToRowAndColumn(i2)

	if (row1+n)%5 >= 0 {
		i1 = table.convertRowAndColumnToIndex((row1+n)%5, col1)
	} else {
		i1 = table.convertRowAndColumnToIndex(5+(row1+n)%5, col1)
	}

	if (row2+n)%5 >= 0 {
		i2 = table.convertRowAndColumnToIndex((row2+n)%5, col2)
	} else {
		i2 = table.convertRowAndColumnToIndex(5+(row2+n)%5, col2)
	}

	return i1, i2
}

func (table *Table) ShiftRectangle(i1, i2 int) (int, int) {
	row1, col1 := table.convertIndexToRowAndColumn(i1)
	row2, col2 := table.convertIndexToRowAndColumn(i2)

	i1 = table.convertRowAndColumnToIndex(row1, col2)
	i2 = table.convertRowAndColumnToIndex(row2, col1)

	return i1, i2
}

func (table *Table) convertIndexToRowAndColumn(index int) (int, int) {
	return index / 5, index % 5
}

func (table *Table) convertRowAndColumnToIndex(row, column int) int {
	return row*5 + column
}

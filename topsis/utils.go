package topsis

func NewFloatVector(n int) []float32 {
	return make([]float32, n)
}

func NewFloatMatrix(m, n int) [][]float32 {
	var matrix [][]float32

	matrix = make([][]float32, m)
	for i := 0; i < m; i++ {
		matrix[i] = NewFloatVector(n)
	}

	return matrix
}

func NewFloatCube(l, m, n int) [][][]float32 {
	var cube[][][]float32

	cube = make([][][]float32, l)
	for i := 0; i < l; i++ {
		cube[i] = NewFloatMatrix(m, n)
	}

	return cube
}

func Sum(elements []float32) float32 {
	sum := float32(0.0)

	for i := range elements {
		sum += elements[i]
	}

	return sum
}

func GetColumn(matrix[][]float32, i int) (column []float32) {
	column = NewFloatVector(len(matrix))

	for j, row := range matrix {
		column[j] = row[i]
	}

	return
}

func GetMin(vector []float32) (element float32) {
	if len(vector) == 0 {
		return
	}

	element = vector[0]

	for i := range vector {
		if element > vector[i] {
			element = vector[i]
		}
	}

	return
}

func GetMax(vector []float32) (element float32) {
	if len(vector) == 0 {
		return
	}

	element = vector[0]

	for i := range vector {
		if element < vector[i] {
			element = vector[i]
		}
	}

	return
}
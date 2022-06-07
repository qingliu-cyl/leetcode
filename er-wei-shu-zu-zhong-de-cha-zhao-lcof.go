package main

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	cow, cel := len(matrix), len(matrix[0])
	for i := 0; i < cow; i++ {
		for j := 0; j < cel; j++ {
			if matrix[i][j] == target {
				return true
			} else if matrix[i][j] > target {
				if j == 0 {
					return false
				}
				break
			}
		}
	}
	return false
}

func main() {

}

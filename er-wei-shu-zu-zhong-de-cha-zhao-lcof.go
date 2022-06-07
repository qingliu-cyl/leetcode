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

//  findNumberIn2DArray 解法2
//  @Description:
//  @param matrix
//  @param target
//  @return bool
func findNumberIn2DArray2(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	cow, cel := 0, len(matrix[0])-1
	for cow < len(matrix) && cel >= 0 {
		if matrix[cow][cel] == target {
			return true
		} else if matrix[cow][cel] > target {
			cel--
		} else {
			cow++
		}
	}

	return false
}

func main() {

}

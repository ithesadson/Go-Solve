package main

import "fmt"

func main() {
	slice := [][]int32{{2, 2}, {2, 3}, {2, 4}, {3, 2},
		{3, 4}, {4, 2}, {4, 3}, {4, 4}, {5, 5}}
	fmt.Println(queensAttack(6, 9, 3, 3, slice))
}

func queensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32 {
	//Number of squares on the bottom, top and sides
	var row_min, row_max, column_min, column_max int32
	row_min = r_q - 1
	row_max = n - r_q
	column_min = c_q - 1
	column_max = n - c_q

	//Number of cross squares
	var top_left, top_right, bottom_left, bottom_right int32
	top_left = min(r_q-1, c_q-1)
	top_right = min(r_q-1, n-c_q)
	bottom_left = min(n-r_q, c_q-1)
	bottom_right = min(n-r_q, n-c_q)

	//Find bottom_leftocked squares
	if obstacles != nil {
		for i := 0; i < len(obstacles); i++ {
			r := obstacles[i][0] // Obstacle Row
			c := obstacles[i][1] // Obstacle Column

			// ---if the queen's row is equal to the obstacle's row
			if r == r_q && c < c_q { // If the same row and queen's column are larger
				row_min = min(row_min, c_q-c-1)
			} else if r == r_q && c > c_q { // If the same row and queen's column are smaller
				row_max = min(row_max, c-c_q-1)

			// ---if the queen's column is equal to the obstacle's column
			} else if c == c_q && r < r_q { // If the same column and row of the queen is larger
				column_min = min(column_min, r_q-r-1)
			} else if c == c_q && r > r_q { // If the same column and queen's row is smaller
				column_max = min(column_max, r-r_q-1)

			// ---Are there any obstacles in the cross
			} else if r-r_q == c-c_q && r < r_q { // r-r_q == c-c_q(Is it on the cross) , Top_Left Checking
				top_left = min(top_left, r_q-r-1)
			} else if r-r_q == c-c_q && r > r_q { // r-r_q == c-c_q(Is it on the cross) , Bottom_right Checking
				bottom_right = min(bottom_right, r-r_q-1)
			} else if r-r_q == c_q-c && r < r_q { // r-r_q == c-c_q(Is it on the cross) , Top_right Checking
				top_right = min(top_right, r_q-r-1)
			} else if r-r_q == c_q-c && r > r_q { // r-r_q == c-c_q(Is it on the cross) , Bottom_left Checking
				bottom_left = min(bottom_left, r-r_q-1)
			}
		}
	}
	total := top_left + top_right + bottom_left + bottom_right + row_min + row_max + column_min + column_max
	return total
}

func min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

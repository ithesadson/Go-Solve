package main

import (
	"fmt"
)

func main() {
	var n, k, rq, cq, r, c int32
	fmt.Scanf("%d %d", &n, &k)   //Scan n and k
	fmt.Scanf("%d %d", &rq, &cq) //Scan r_q and c_q

	obstacles := [][]int32{}
	for i := 0; i < int(k); i++ {
		fmt.Scanf("%d %d", &r, &c) //Scan obstacles
		obstacles = append(obstacles, []int32{r, c})
	}
	fmt.Println(queensAttack(n, k, rq, cq, obstacles))
}

func queensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32 {
	//Number of squares on the bottom, top and sides
	var rmin, rmax, cmin, cmax int32
	rmin = r_q - 1
	rmax = n - r_q
	cmin = c_q - 1
	cmax = n - c_q

	//Number of cross squares
	var tl, tr, bl, br int32
	tl = min(r_q-1, c_q-1)
	tr = min(r_q-1, n-c_q)
	bl = min(n-r_q, c_q-1)
	br = min(n-r_q, n-c_q)

	//Find blocked squares
	if obstacles != nil {
		for i := 0; i < len(obstacles); i++ {
			r := obstacles[i][0]
			c := obstacles[i][1]
			if r == r_q && c < c_q {
				rmin = min(rmin, c_q-c-1)
			} else if r == r_q && c > c_q {
				rmax = min(rmax, c-c_q-1)
			} else if c == c_q && r < r_q {
				cmin = min(cmin, r_q-r-1)
			} else if c == c_q && r > r_q {
				cmax = min(cmax, r-r_q-1)
			} else if r-r_q == c-c_q && r < r_q {
				tl = min(tl, r_q-r-1)
			} else if r-r_q == c-c_q && r > r_q {
				br = min(br, r-r_q-1)
			} else if r-r_q == c_q-c && r < r_q {
				tr = min(tr, r_q-r-1)
			} else if r-r_q == c_q-c && r > r_q {
				bl = min(bl, r-r_q-1)
			}
		}
	}
	total := tl + tr + bl + br + rmin + rmax + cmin + cmax
	return total
}

func min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	//Test case
	player := []int32{70, 80, 105}
	ranked := []int32{100, 90, 90, 80}
	fmt.Println(climbingLeaderboard(ranked, player))

	//Test case
	player1 := []int32{5, 25, 50, 120}
	ranked1 := []int32{100, 100, 50, 40, 40, 20, 10}
	fmt.Println(climbingLeaderboard(ranked1, player1))

}

func climbingLeaderboard(ranked []int32, player []int32) []int32 {
	newleaderboard := ranked[:1]
	temp := newleaderboard[0]

	//Remove repetitive elements from the array
	for _, x := range ranked[1:] {
		if x != temp {
			newleaderboard = append(newleaderboard, x)
		}
		temp = x
	}

	ret_arr := make([]int32, 0, len(player)) //Define return array
	for _, y := range player {
		//Compare the player score with the leaderboard
		//max_score returns its place in the ranking
		max_score := sort.Search(
			len(newleaderboard), func(i int) bool { return newleaderboard[i] <= y })
		ret_arr = append(ret_arr, int32(max_score+1))
	}
	return ret_arr
}

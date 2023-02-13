package main

import "fmt"

func main() {
	// This question part of Nim game Strategy

	// For Default test-case | Sample Input 2 3
	//  a | 2 = 10 (binary)
	//  b | 3 = 11 (binary) --> a+b = binary(01) = 1 | allxor=1

	arr := [2]int32{2, 3}
	var win_counter int32 = 0
	var allxor int32 = 0
	N := len(arr)

	for i := 0; i < N; i++ {
		allxor ^= arr[i]
	}

	for j := 0; j < N; j++ {
		//For player 1 to win, the array element must be greater than element * allxor.
		if arr[j] > (allxor ^ arr[j]) {
			win_counter += 1
		}
	}
	fmt.Println(win_counter)
}

/* Function solution
func chocolateInBox(arr []int32) int32 {
	// This question part of Nim game Strategy
	var win_counter int32 = 0
	var allxor int32 = 0
	N := len(arr)

	for i := 0; i < N; i++ {
		allxor ^= arr[i]
	}
	for i := 0; i < N; i++ {
		if arr[i] > (allxor ^ arr[i]) {
			win_counter += 1
		}
	}
	return win_counter
}
*/

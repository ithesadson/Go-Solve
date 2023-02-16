package main

import "fmt"

func main() {
	var n int
	//How many elements will the array have
	fmt.Scan(&n)
	//Define an array of "n" elements
	arr := make([]int32, n)

	for i := 0; i < n; i++ {
		//Enter the elements of the array
		fmt.Scan(&arr[i])
	}
	fmt.Println("Sample Output:", chocolateInBox(arr))
}

func chocolateInBox(arr []int32) int32 {
	//This function solution includes a nim game strategy
	var win_counter int32 = 0
	var allxor int32 = 0 // XOR logic operator
	N := len(arr)
	for i := 0; i < N; i++ {
		//Sum of array elements with XOR Gate
		allxor ^= arr[i]
	}
	for i := 0; i < N; i++ {
		//Calculation of probabilities leading to a decisive win.
		//Applying and comparing the xor gate for all elements.
		if arr[i] > (allxor ^ arr[i]) {
			win_counter += 1
		}
	}
	return win_counter
}

/* Test Case ->
Input (stdin)
2
2 3

Expected Output
1
---------------
Input (stdin)
5
251149 86127 711523 501067 617190

Expected Output
1
*/

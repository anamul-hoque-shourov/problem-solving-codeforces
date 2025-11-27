package ICPC_Assiut_Winter_Camp_Newcomers_20

import "fmt"

func solution1(n, x, y int) int {
	profit := n*y - n*x
	return profit
}

func TestCode1() {
	result := solution1(2, 3, 5)
	fmt.Println(result)
}

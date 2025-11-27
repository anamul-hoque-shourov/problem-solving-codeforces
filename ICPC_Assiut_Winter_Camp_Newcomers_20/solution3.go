package ICPC_Assiut_Winter_Camp_Newcomers_20

import "fmt"

func solution3() {
	var n int
	fmt.Scan(&n)

	if n == 10 {
		fmt.Println(0)
		return
	}

	coins := []int{25, 10, 5, 1}
	count := 0

	for _, coin := range coins {
		count = count + n/coin
		n = n % coin
	}

	fmt.Println(count)
}

func TestCode3() {
	solution3()
}

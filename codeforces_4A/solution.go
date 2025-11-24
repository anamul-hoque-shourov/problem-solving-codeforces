package codeforces_4A

import "fmt"

func Solution() {
	var w int
	_, err := fmt.Scanf("%d", &w)
	if err != nil {
		return
	}

	if w > 2 && w%2 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func TestCode() {
	Solution()
}

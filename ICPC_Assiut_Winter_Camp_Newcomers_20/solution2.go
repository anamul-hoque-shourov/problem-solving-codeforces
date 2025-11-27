package ICPC_Assiut_Winter_Camp_Newcomers_20

import (
	"fmt"
)

func solution2() {
	var a, b, c float64
	fmt.Scanf("%f%f%f", &a, &b, &c)
	result := (2*a + 3*b) * 5 * c
	fmt.Printf("%.6f\n", result)
}

func TestCode2() {
	solution2()
}

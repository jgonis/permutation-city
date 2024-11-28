package permutations

import "fmt"

func GeneratePermutations(n uint) {
	blah := []int{}
	for i := range n {
		blah = append(blah, int(i))
	}
	generate(n, blah)
}

func generate(k uint, A []int) {
	if k == 1 {
		fmt.Println(A)
		return
	} else {
		generate(k-1, A)

		for i := range k - 1 {
			if k%2 == 0 {
				swap(i, k-1, A)
			} else {
				swap(0, k-1, A)
			}
			generate(k-1, A)
		}
	}
}

func swap(a, b uint, A []int) {
	temp := A[a]
	A[a] = A[b]
	A[b] = temp
}

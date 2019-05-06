package main

import "fmt"

func main() {
	var a []int
	a = []int{1, 2, 3}
	a[1] = 10
	fmt.Println(a)

	s := make([]int, 5, 10)
	println(len(s), cap(s))

	n_s := []int{0, 1, 2, 3, 4, 5}
	n_s = n_s[2:5]
	n_s = n_s[1:]
	fmt.Println(n_s)

	var empty_s []int
	if empty_s == nil {
		println("Nil Slice")
	}
	println(len(empty_s), cap(empty_s))

	a_s := []int{0, 1}
	a_s = append(a_s, 2)
	a_s = append(a_s, 3, 4, 5)

	fmt.Println(a_s)

	// 뭐야 원래 camel 이었어?? snake 인줄 알았는데
	sliceA := make([]int, 0, 3)

	for i := 1; i <= 15; i++ {
		sliceA = append(sliceA, i)
		fmt.Println(len(sliceA), cap(sliceA))
	}

	fmt.Println(sliceA)

	sA := []int{1, 2, 3}
	sB := []int{4, 5, 6}

	sA = append(sA, sB...)
	fmt.Println(sA)

	source := []int{0, 1, 2}
	target := make([]int, len(source), cap(source)*2)
	copy(target, source)
	fmt.Println(target)
	println(len(target), cap(target))
}

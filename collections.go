package main

func main() {
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	println(a[1])

	var a1 = [3]int{1, 2, 3}
	var a3 = [...]int{1, 2, 3}

	println(a1[2])
	println(a3[2])

	var multiArray [3][4][5]int
	multiArray[0][1][2] = 10

	var multiArray2 = [2][3]int{
		{1, 2, 3},
		{4, 5, 6}, // 끝에 콤마 추가 해야함
	}
	println(multiArray2[1][2])
}

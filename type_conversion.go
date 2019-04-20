package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	var i int = 100
	var u uint = uint(i)
	var f float32 = float32(i)
	println(f, u)

	str := "ABC"
	bytes := []byte(str)
	str2 := string(bytes)
	println(bytes, str2)

	const a uint32 = 3
	const b uint32 = 1

	println(a & b)
	println(a | b)
	println(a >> 1)

	v := make([]byte, 8)
	binary.BigEndian.PutUint32(v, a)
	fmt.Println(v)

	println(&v)
	println(*(&str))
}

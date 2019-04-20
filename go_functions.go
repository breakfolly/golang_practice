package main

func main() {
	msg := "Hello"
	say(&msg)
	println(msg)

	say1("This", "is", "a", "book")
	say1("Hi")

	count, total := sum(1, 7, 3, 5, 9)
	println(count, total)
}

func say(msg *string) {
	println(*msg)
	*msg = "Changed"
}

func say1(msg ...string) {
	for _, s := range msg {
		println(s)
	}
}

func sum(nums ...int) (int, int) {
	s := 0
	count := 0
	for _, n := range nums {
		s += n
		count++
	}
	return count, s
}

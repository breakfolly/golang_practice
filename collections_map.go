package main

import "fmt"

func main() {
	var m map[int]string

	m = make(map[int]string)

	m[901] = "Apple"
	m[134] = "Grape"
	m[777] = "Tomato"

	str := m[134]
	println(str)

	noData := m[999]
	println(noData)

	delete(m, 777)

	tickers := map[string]string{
		"GOOG": "Google Inc",
		"MSFT": "Microsoft",
		"FB":   "FaceBook",
		"AMZS": "Amazon",
	}

	val, exists := tickers["MSFT"]
	println(val)

	if !exists {
		println("No MSFT ticker")
	}

	myMap := map[string]string{
		"A": "Apple",
		"B": "Banana",
		"C": "Charlie",
	}

	for key, val := range myMap {
		fmt.Println(key, val)
	}
}

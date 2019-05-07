package main

import "custom_package"

func main() {
	song := custom_package.GetMusic("Alicia Keys")
	println(song)
}

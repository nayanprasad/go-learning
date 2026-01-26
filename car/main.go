package main

func main() {
	cars := cars{"white car", "red car"}

	cars = append(cars, "blue car")

	cars.print()
}

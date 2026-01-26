package main

// declare a custom type  like class members
type cars []string

// this funciton is called receiver fucntion and c here is act as "this"
func (c cars) print() {
	for _, car := range c {
		println(car)
	}
}

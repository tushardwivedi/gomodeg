package gomodeg

import "github.com/tushardwivedi/gomod1"

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func Meoh() string {
	return "Meoh!"
}

func BigBark() string {
	return gomod1.WhenGrownUps(Bark())
}

func BigBarks() string {
	return gomod1.WhenGrownUps(Meoh())
}

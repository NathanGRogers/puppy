package puppy

import (
	"fmt"

	"github.com/NathanGRogers/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())

}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())

}

func From11() {
	fmt.Println("I'm from version version 1.1.1")

}

func From12() {
	fmt.Println("I'm from version version 1.2.0")

}

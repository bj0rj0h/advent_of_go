package main

import "fmt"

var heading string

var l = Location{0, 0}

type Location struct {
	X int
	Y int
}

func move(steps int) {

	switch heading {
	case "N":
		l.Y = l.Y + steps
	case "E":
		l.X = l.X + steps
	case "S":
		l.Y = l.Y - steps
	case "W":
		l.X = l.X - steps
	}
}

/*func turn(t string){

	switch t{
		case ""
	}
}*/

func main() {

	heading = "N"
	move(4)
	fmt.Println("Location is: ", l)
}

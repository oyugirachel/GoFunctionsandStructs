package main
import (
	"fmt"
)
type Bootcamp struct {
	Lat float64
	Lon float64
}
type User struct {
	Id       int
	Name     string
	Location string
}

type Player struct {
	Id       int
	Name     string
	Location string
	GameId	 int
}
func main() {
	x := new(Bootcamp)
	y := &Bootcamp{}
	fmt.Println(*x == *y)
	p := Player{}
	p.Id = 42
	p.Name = "Matt"
	p.Location = "LA"
	p.GameId = 90404
	fmt.Printf("%+v", p)
}


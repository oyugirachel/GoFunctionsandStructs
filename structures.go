package main

import "fmt"

type Behave struct {
	Id             int
	Name, Location string
}
type Exercise struct {
	Behave
	GameId int
}

func main() {
	p := Exercise{}
	p.Id = 42
	p.Name = "Matt"
	p.Location = "LA"
	p.GameId = 90404
	fmt.Printf("%+v", p)
}
